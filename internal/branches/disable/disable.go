package disable

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-errors/errors"
	"github.com/spf13/afero"
	"github.com/hk8xb/gentabase-cli/internal/utils"
	"github.com/hk8xb/gentabase-cli/internal/utils/flags"
)

func Run(ctx context.Context, fsys afero.Fs) error {
	resp, err := utils.GetGentabaseAPI().V1DisablePreviewBranchingWithResponse(ctx, flags.ProjectRef)
	if err != nil {
		return errors.Errorf("failed to disable preview branching: %w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.Errorf("unexpected disable branching status %d: %s", resp.StatusCode(), string(resp.Body))
	}
	fmt.Println("Disabled preview branching for project:", flags.ProjectRef)
	return nil
}
