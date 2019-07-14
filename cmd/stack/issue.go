package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v26/github"
	"github.com/jlucktay/stack/pkg/common"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// issue flow:
// 0. get current stack directory
// 1. send issue to GitHub with appropriate directory tag

func createIssue(text ...string) {
	fmt.Printf("your text: %#v\n", text)

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

	fmt.Printf("stackPath: '%s'\n", stackPath)

	// 1
	// WIP

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: viper.GetString("github.pat"),
		},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	issueRequest := &github.IssueRequest{
		Title: github.String("hello world"),
		Labels: &[]string{
			stackPath,
		},
	}

	// issue, response, errCreate := client.Issues.Create(
	// 	ctx,
	// 	viper.GetString("github.org"),
	// 	viper.GetString("github.repo"),
	// 	issueRequest,
	// )
	issue, response, errCreate := client.Issues.Create(ctx, "jlucktay", "stack", issueRequest)
	if errCreate != nil {
		log.Fatal(errors.Wrap(errCreate, "errCreate!\n"))
	}

	fmt.Printf("\nissue:\n\t'%s'\n", issue)
	fmt.Printf("\nresponse:\n\t'%#v'\n", *response)
}
