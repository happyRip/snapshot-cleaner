package main

import (
	"fmt"
	"log"
	"time"

	"github.com/docker/hub-tool/pkg/hub"

	"github.com/happyRip/snapshot-cleaner/pkg/env"
)

const day = 24 * time.Hour

func main() {
	env, err := env.New()
	if err != nil {
		log.Fatalf("Failed to initialize the environment | %s", err)
	}

	hubClient, err := hub.NewClient(hub.WithHubAccount(env.Username), hub.WithPassword(env.Password))
	if err != nil {
		log.Fatalf("Can't initiate hubClient | %s", err.Error())
	}

	tags, _, err := hubClient.GetTags(env.Repository)
	if err != nil {
		log.Fatalf("Can't get tags list | %s", err.Error())
	}

	for _, tag := range tags {
		if tag.LastUpdated.Before(time.Now().Add(-env.Retention * day)) {
			fmt.Println(tag.Name)
		}
	}
}
