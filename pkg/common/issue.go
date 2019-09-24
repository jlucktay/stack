package common

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-github/v27/github"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// issue flow:
// 0. get current stack directory
// 0.1. get current GitHub username
// 1. send issue to GitHub with appropriate directory tag
// 2. print the URL of the newly-created issue

func CreateIssue(title string, text ...string) {
	// 0
	stackPath := mustGetStackPath()

	// set up GitHub auth
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: viper.GetString("github.pat"),
		},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// 0.1
	assignee := mustGetCurrentGitHubLogin(client)

	// 1
	issueRequest := &github.IssueRequest{
		Title: &title,
		Body:  github.String(strings.Join(text, " ")),
		Labels: &[]string{
			stackPath,
		},
		Assignee: &assignee,
	}

	issue, _, errCreate := client.Issues.Create(
		ctx,
		viper.GetString("github.org"),
		viper.GetString("github.repo"),
		issueRequest,
	)
	if errCreate != nil {
		panic(errors.Wrap(errCreate, "errCreate!\n"))
	}

	// 2
	newIssue, errParse := url.Parse(issue.GetHTMLURL())
	if errParse != nil {
		panic(errors.Wrap(errParse, "errParse!\n"))
	}

	fmt.Printf("New issue: %s\n", newIssue)
}
