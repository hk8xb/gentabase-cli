package main

import (
	"context"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hk8xb/gentabase-cli/pkg/api"
	"github.com/hk8xb/gentabase-cli/pkg/config"
	"github.com/hk8xb/gentabase-cli/pkg/function"
)

func main() {
	if err := deploy(context.Background(), os.DirFS(".")); err != nil {
		log.Fatalln(err)
	}
}

func deploy(ctx context.Context, fsys fs.FS) error {
	project := os.Getenv("GENTABASE_PROJECT_ID")
	apiClient := newAPIClient(os.Getenv("GENTABASE_ACCESS_TOKEN"))
	functionClient := function.NewEdgeRuntimeAPI(project, apiClient)
	fc := config.FunctionConfig{"my-slug": {
		Entrypoint: "gentabase/functions/my-slug/index.ts",
		ImportMap:  "gentabase/functions/import_map.json",
	}}
	return functionClient.Deploy(ctx, fc, fsys)
}

func newAPIClient(token string) api.ClientWithResponses {
	header := func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	}
	client := api.ClientWithResponses{ClientInterface: &api.Client{
		// Ensure the server URL always has a trailing slash
		Server: "https://gentabase.dev/management/",
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
		RequestEditors: []api.RequestEditorFn{header},
	}}
	return client
}
