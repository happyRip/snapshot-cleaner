package hub

import (
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/docker/hub-tool/pkg/hub"

	"github.com/happyRip/snapshot-cleaner/pkg/env"
)

type Client struct {
	hubClient *hub.Client
	env       *env.Environment
}

func NewClient() (*Client, error) {
	env, err := env.New()
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize the environment")
	}
	c, err := hub.NewClient(
		hub.WithHubAccount(env.Username),
		hub.WithPassword(env.Password),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize the client")
	}
	return &Client{
		hubClient: c,
		env:       env,
	}, nil
}

func (c Client) ListTags() ([]hub.Tag, error) {
	tags, _, err := c.hubClient.GetTags(c.env.Repository)
	return tags, errors.Wrap(err, fmt.Sprintf("failed to list tags for repository %q", c.env.Repository))
}

func (c Client) ListTagsBefore(t time.Time) ([]hub.Tag, error) {
	tags, err := c.ListTags()
	if err != nil {
		return nil, errors.Wrap(err, "failed to list tags")
	}
	for i, tag := range tags {
		if tag.LastUpdated.Before(t) {
			if i == 0 {
				return nil, nil
			}
			return tags[0:i], nil
		}
	}
	return tags, nil
}

func (c Client) CleanupTags() error {
	tags, err := c.ListTags()
	if err != nil {
		return errors.Wrap(err, "failed to list tags")
	}
	repository := c.env.Repository
	for _, tag := range tags {
		if tag.LastUpdated.Before(time.Now().Add(-c.env.Retention)) {
			name := tag.Name
			if err := c.hubClient.RemoveTag(repository, name); err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to remove tag %q", name))
			}
		}
	}
	return nil
}

func (c Client) HubClient() *hub.Client {
	return c.hubClient
}

func (c Client) Env() *env.Environment {
	return c.env
}
