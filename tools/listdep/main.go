package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/hk8xb/gentabase-cli/pkg/config"
)

func main() {
	const hk8xbPrefix = "ghcr.io/hk8xb/"
	external := make([]string, 0)
	for _, img := range config.Images.Services() {
		if !strings.HasPrefix(img, hk8xbPrefix) ||
			strings.Contains(img, "/logflare:") {
			external = append(external, img)
		}
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(external); err != nil {
		log.Fatal(err)
	}
}
