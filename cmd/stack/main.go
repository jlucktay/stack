package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Take ldflags from GoReleaser
//nolint
var (
	version, commit, builtBy string
	date                     = time.Now().UTC().String()
)

func main() {
	viper.SetConfigName("stack.config") // name of config file (without extension)
	viper.SetConfigType("json")
	viper.AddConfigPath("$HOME/.config/stack") // path to look for the config file in
	viper.AddConfigPath(".")                   // optionally look for config in the working directory

	errViperRead := viper.ReadInConfig() // Find and read the config file
	if errViperRead != nil {             // Handle errors reading the config file
		log.Fatalf("fatal error with config file: %s\n", errViperRead)
	}

	// Define Usage() for flags and '--help'
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Available commands for '%s':\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  build    Queue a plan to build this Terraform stack\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  cancel   Cancel release(s) of built/planned Terraform stack\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  destroy  Queue a plan to destroy this Terraform stack\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  init     Initialise this Terraform stack against remote state\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  issue    Add/update a GitHub issue for this Terraform stack\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  version  Show details of this binary's current version\n")
		flag.PrintDefaults()
	}

	// Parse out current git branch for use as a default for one of the 'build' flags
	gitRaw, errExec := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if errExec != nil {
		log.Fatal(errExec)
	}
	currentBranch := strings.TrimSpace(string(gitRaw))

	// Set up 'build' subcommand
	buildCommand := flag.NewFlagSet("build", flag.ExitOnError)
	buildBranch := buildCommand.String("branch", currentBranch, "If given, plan from this branch.\n"+
		"Defaults to the current branch.")
	buildTargets := buildCommand.String("target", "", "If given, target these specific Terraform resources only.\n"+
		"Delimit multiple target IDs with a semi-colon ';'.")

	// Set up 'cancel' subcommand
	cancelCommand := flag.NewFlagSet("cancel", flag.ExitOnError)

	// Set up 'destroy' subcommand
	destroyCommand := flag.NewFlagSet("destroy", flag.ExitOnError)
	destroyBranch := destroyCommand.String("branch", currentBranch, "If given, plan from this branch.\n"+
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

	// Set up 'version' subcommand
	versionCommand := flag.NewFlagSet("version", flag.ExitOnError)
	versionCommand.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"stack v%s from commit %s, built %s by %s.\n",
			version, commit, date, builtBy)
	}

	// Check which subcommand was given, and parse accordingly
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case buildCommand.Name():
		// Parse and execute 'build' subcommand
		if errParse := buildCommand.Parse(os.Args[2:]); errParse != nil {
			log.Fatalf("error parsing build flags: %v", errParse)
		}
		stackQueue(*buildBranch, *buildTargets, uint(viper.GetInt("azureDevOps.buildDefID")))
	case cancelCommand.Name():
		if errParse := cancelCommand.Parse(os.Args[2:]); errParse != nil {
			log.Fatalf("error parsing cancel flags: %v", errParse)
		}
	case destroyCommand.Name():
		// Parse and execute 'destroy' subcommand
		if errParse := destroyCommand.Parse(os.Args[2:]); errParse != nil {
			log.Fatalf("error parsing destroy flags: %v", errParse)
		}
		stackQueue(*destroyBranch, *destroyTargets, uint(viper.GetInt("azureDevOps.destroyDefID")))
	case initCommand.Name():
		// Execute on 'init' subcommand
		initStack()
	case issueCommand.Name():
		// Execute on 'issue' subcommand
		if len(os.Args[2:]) == 0 {
			issueCommand.Usage()
			log.Fatalf("No issue text was given!")
		}
		createIssue(os.Args[2:]...)
	case versionCommand.Name():
		versionCommand.Usage()
		os.Exit(0)
	default:
		fmt.Printf("'%v' is not a valid command.\n", os.Args[1])
		flag.Usage()
		os.Exit(1)
	}

	// Execute on 'cancel' subcommand
	if cancelCommand.Parsed() {
		fmt.Println("'cancel' is not yet implemented.") // TODO
		os.Exit(0)

		if cancelCommand.NFlag() == 0 {
			fmt.Println("Please specify one or more releases to cancel.")
			cancelCommand.PrintDefaults()
			os.Exit(1)
		}

		cancelCommand.Visit(func(f *flag.Flag) {})
	}
}
