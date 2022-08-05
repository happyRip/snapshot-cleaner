package main

import (
	"log"

	"github.com/pkg/errors"

	"github.com/happyRip/snapshot-cleaner/cmd"
	"github.com/happyRip/snapshot-cleaner/pkg/hub"
)

func main() {
	_, err := hub.NewClient()
	if err != nil {
		log.Fatalf("failed to initialize the hub client: %+v\n", errors.Unwrap(err))
	}

	cmd.Execute()
}
