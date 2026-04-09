package get

import (
	"context"
	"fmt"
	"os"

	"github.com/go-errors/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/spf13/afero"
	"github.com/hk8xb/gentabase-cli/internal/projects/apiKeys"
	"github.com/hk8xb/gentabase-cli/internal/utils"
	"github.com/hk8xb/gentabase-cli/internal/utils/flags"
	"github.com/hk8xb/gentabase-cli/pkg/api"
	"github.com/hk8xb/gentabase-cli/pkg/cast"
)

func Run(ctx context.Context, branchId string, fsys afero.Fs) error {
	detail, err := getBranchDetail(ctx, branchId)
	if err != nil {
		return err
	}

	if utils.OutputFormat.Value != utils.OutputPretty {
		keys, err := apiKeys.RunGetApiKeys(ctx, detail.Ref)
		if err != nil {
			return err
		}
		pooler, err := utils.GetPoolerConfigPrimary(ctx, detail.Ref)
		if err != nil {
			return err
		}
		envs := toStandardEnvs(detail, pooler, keys)
		return utils.EncodeOutput(utils.OutputFormat.Value, os.Stdout, envs)
	}

	table := `|HOST|PORT|USER|PASSWORD|JWT SECRET|POSTGRES VERSION|STATUS|
|-|-|-|-|-|-|-|
` + fmt.Sprintf(
		"|`%s`|`%d`|`%s`|`%s`|`%s`|`%s`|`%s`|\n",
		detail.DbHost,
		detail.DbPort,
		*detail.DbUser,
		*detail.DbPass,
		*detail.JwtSecret,
		detail.PostgresVersion,
		detail.Status,
	)

	return utils.RenderTable(table)
}

func getBranchDetail(ctx context.Context, branchId string) (api.BranchDetailResponse, error) {
	var result api.BranchDetailResponse
	if err := uuid.Validate(branchId); err != nil && !utils.ProjectRefPattern.Match([]byte(branchId)) {
		resp, err := utils.GetGentabaseAPI().V1GetABranchWithResponse(ctx, flags.ProjectRef, branchId)
		if err != nil {
			return result, errors.Errorf("failed to find branch: %w", err)
		} else if resp.JSON200 == nil {
			return result, errors.Errorf("unexpected find branch status %d: %s", resp.StatusCode(), string(resp.Body))
		}
		branchId = resp.JSON200.ProjectRef
	}
	resp, err := utils.GetGentabaseAPI().V1GetABranchConfigWithResponse(ctx, branchId)
	if err != nil {
		return result, errors.Errorf("failed to get branch: %w", err)
	} else if resp.JSON200 == nil {
		return result, errors.Errorf("unexpected get branch status %d: %s", resp.StatusCode(), string(resp.Body))
	}
	masked := "******"
	if resp.JSON200.DbUser == nil {
		resp.JSON200.DbUser = &masked
	}
	if resp.JSON200.DbPass == nil {
		resp.JSON200.DbPass = &masked
	}
	if resp.JSON200.JwtSecret == nil {
		resp.JSON200.JwtSecret = &masked
	}
	return *resp.JSON200, nil
}

func toStandardEnvs(detail api.BranchDetailResponse, pooler api.SupavisorConfigResponse, keys []api.ApiKeyResponse) map[string]string {
	direct := pgconn.Config{
		Host:     detail.DbHost,
		Port:     cast.UIntToUInt16(cast.IntToUint(detail.DbPort)),
		User:     *detail.DbUser,
		Password: *detail.DbPass,
		Database: "postgres",
	}
	config, err := utils.ParsePoolerURL(pooler.ConnectionString)
	if err != nil {
		fmt.Fprintln(os.Stderr, utils.Yellow("WARNING:"), err)
		config = &direct
	} else {
		config.Password = direct.Password
	}
	envs := apiKeys.ToEnv(keys)
	envs["POSTGRES_URL"] = utils.ToPostgresURL(*config)
	envs["POSTGRES_URL_NON_POOLING"] = utils.ToPostgresURL(direct)
	envs["GENTABASE_URL"] = "https://" + utils.GetGentabaseHost(detail.Ref)
	envs["GENTABASE_JWT_SECRET"] = *detail.JwtSecret
	return envs
}
