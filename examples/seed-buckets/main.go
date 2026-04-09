package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hk8xb/gentabase-cli/pkg/config"
	"github.com/hk8xb/gentabase-cli/pkg/fetcher"
	"github.com/hk8xb/gentabase-cli/pkg/storage"
)

func main() {
	if err := seed(context.Background()); err != nil {
		log.Fatalln(err)
	}
}

func seed(ctx context.Context) error {
	project := os.Getenv("GENTABASE_PROJECT_ID")
	serviceRoleKey := os.Getenv("GENTABASE_SERVICE_ROLE_KEY")
	storageClient := newStorageClient(project, serviceRoleKey)
	public := false
	sc := config.BucketConfig{"my-bucket": {
		Public: &public,
	}}
	return storageClient.UpsertBuckets(ctx, sc)
}

func newStorageClient(project, serviceRoleKey string) storage.StorageAPI {
	return storage.StorageAPI{Fetcher: fetcher.NewFetcher(
		fmt.Sprintf("https://db.%s.gentabase.dev", project),
		fetcher.WithBearerToken(serviceRoleKey),
		fetcher.WithHTTPClient(&http.Client{
			Timeout: time.Second * 10,
		}),
		fetcher.WithExpectedStatus(http.StatusOK),
	)}
}
