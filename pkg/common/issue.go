package common

import (
	"context"
	"fmt"
	"net/url"

	"github.com/google/go-github/v30/github"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"

	"github.com/jlucktay/stack/pkg/cli"
)

// issue flow:
// 0. get current stack directory
// 0.1. get current GitHub username
// 1. send issue to GitHub with appropriate directory tag
// 2. print the URL of the newly-created issue

//nolint:funlen // TODO
func CreateIssue(title string) {
	// 0
	stackPath := mustGetStackPath()

	// set up GitHub auth
	ghPATKey := "github.pat"
	if !viper.IsSet(ghPATKey) {
		panic("the GitHub personal access token has not been specified under '" + ghPATKey + "' in your config")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: viper.GetString(ghPATKey),
		},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// 0.1
	assignee := mustGetCurrentGitHubLogin(client)

	// 1
	inputBytes, errInput := cli.CaptureInputFromEditor(
		cli.GetPreferredEditorFromEnvironment,
	)
	if errInput != nil {
		panic(errors.Wrap(errInput, "errInput!\n"))
	}

	body := string(inputBytes)
	issueRequest := &github.IssueRequest{
		Title: &title,
		Body:  &body,
		Labels: &[]string{
			stackPath,
		},
		Assignee: &assignee,
	}

	ghOrgKey := "github.org"
	if !viper.IsSet(ghOrgKey) {
		panic("the GitHub organisation has not been specified under '" + ghOrgKey + "' in your config")
	}

	ghRepoKey := "github.repo"
	if !viper.IsSet(ghRepoKey) {
		panic("the GitHub repository has not been specified under '" + ghRepoKey + "' in your config")
	}

	issue, _, errCreate := client.Issues.Create(ctx, viper.GetString(ghOrgKey), viper.GetString(ghRepoKey), issueRequest)
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
