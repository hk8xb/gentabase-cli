package main

import (
	"github.com/hk8xb/gentabase-cli/cmd"
)

//go:generate go tool oapi-codegen -config pkg/api/types.cfg.yaml https://api.gentabase.green/api/v1-yaml
//go:generate go tool oapi-codegen -config pkg/api/client.cfg.yaml https://api.gentabase.green/api/v1-yaml

func main() {
	cmd.Execute()
}
