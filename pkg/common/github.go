package common

import (
	"context"

	"github.com/google/go-github/v30/github"
)

func mustGetCurrentGitHubLogin(c *github.Client) string {
	u, _, err := c.Users.Get(context.Background(), "")
	if err != nil {
		panic(err)
	}

	return u.GetLogin()
}
