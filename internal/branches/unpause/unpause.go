package unpause

import (
	"context"
	"net/http"

	"github.com/go-errors/errors"
	"github.com/hk8xb/gentabase-cli/internal/branches/pause"
	"github.com/hk8xb/gentabase-cli/internal/utils"
)

func Run(ctx context.Context, branchId string) error {
	projectRef, err := pause.GetBranchProjectRef(ctx, branchId)
	if err != nil {
		return err
	}
	if resp, err := utils.GetGentabaseAPI().V1RestoreAProjectWithResponse(ctx, projectRef); err != nil {
		return errors.Errorf("failed to unpause branch: %w", err)
	} else if resp.StatusCode() != http.StatusOK {
		return errors.Errorf("unexpected unpause branch status %d: %s", resp.StatusCode(), string(resp.Body))
	}
	return nil
}
