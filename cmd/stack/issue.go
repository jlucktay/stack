package main

import (
	"fmt"
	"log"

	"github.com/jlucktay/stack/pkg/common"
	"github.com/spf13/viper"
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
	// TODO
}
