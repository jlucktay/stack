package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

//nolint
func main() {
	// Set up 'destroy' subcommand
	destroyCommand := flag.NewFlagSet("destroy", flag.ExitOnError)
	destroyBranch := destroyCommand.String("branch", "", "If given, plan from this branch.\n"+
		"Defaults to the current branch.")
	destroyTargets := destroyCommand.String("target", "", "If given, target these specific Terraform resources only.\n"+
		"Delimit multiple target IDs with a semi-colon ';'.")

	// Set up 'init' subcommand
	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	initCommand.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"no flags here, just stacks to initialise!\n")
	}

	// Set up 'issue' subcommand
	issueCommand := flag.NewFlagSet("issue", flag.ExitOnError)
	issueCommand.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "issueCommand.Usage: just give me some words!\n")
	}

	// Check which subcommand was given, and parse accordingly
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case destroyCommand.Name():
		// Parse and execute 'destroy' subcommand
		if errParse := destroyCommand.Parse(os.Args[2:]); errParse != nil {
			panic(fmt.Sprintf("error parsing destroy flags: %v", errParse))
		}
		stackQueue(*destroyBranch, *destroyTargets, uint(viper.GetInt("azureDevOps.destroyDefID")))
	case initCommand.Name():
		// Execute on 'init' subcommand
		initStack()
	case issueCommand.Name():
		// Execute on 'issue' subcommand
		if len(os.Args[2:]) == 0 {
			issueCommand.Usage()
			panic(fmt.Sprintf("No issue text was given!"))
		}
		createIssue(os.Args[2:]...)
	default:
		fmt.Printf("'%v' is not a valid command.\n", os.Args[1])
		flag.Usage()
		os.Exit(1)
	}
}
