package storage

import "github.com/hk8xb/gentabase-cli/pkg/fetcher"

type StorageAPI struct {
	*fetcher.Fetcher
}

const PAGE_LIMIT = 100
