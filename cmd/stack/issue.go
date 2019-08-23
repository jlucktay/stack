package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/google/go-github/v27/github"
	"github.com/jlucktay/stack/pkg/common"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// issue flow:
// 0. get current stack directory
// 1. send issue to GitHub with appropriate directory tag
// 2. print the URL of the newly-created issue

func createIssue(text ...string) {
	// 0
	stackPath, errStackPath := common.GetStackPath(
		viper.GetString("stackPrefix"),
		fmt.Sprintf(
			"github.com/%s/%s",
			viper.GetString("github.org"),
			viper.GetString("github.repo"),
		),
	)
	if errStackPath != nil {
		log.Fatal(errStackPath)
	}

	// 1
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: viper.GetString("github.pat"),
		},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	issueRequest := &github.IssueRequest{
		Title: github.String(strings.Join(text, " ")),
		Labels: &[]string{
			stackPath,
		},
	}

	issue, _, errCreate := client.Issues.Create(
		ctx,
		viper.GetString("github.org"),
		viper.GetString("github.repo"),
		issueRequest,
	)
	if errCreate != nil {
		log.Fatal(errors.Wrap(errCreate, "errCreate!\n"))
	}

	// 2
	newIssue, errParse := url.Parse(issue.GetHTMLURL())
	if errParse != nil {
		log.Fatal(errors.Wrap(errParse, "errParse!\n"))
	}

	fmt.Printf("New issue: %s\n", newIssue)
}
