package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
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
		fmt.Fprintf(flag.CommandLine.Output(), "  build    Queue a build for a Terraform stack\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  cancel   Cancel release(s) of a built/planned Terraform stack\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  issue    Add/update an issue for a Terraform stack\n")
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
	buildBranch := buildCommand.String("branch", currentBranch, "If given, build from this branch.\n"+
		"Defaults to the current branch.")
	buildTargets := buildCommand.String("target", "", "If given, target these specific Terraform resources only.\n"+
		"Delimit between targets with a semi-colon ';'.")

	// Set up 'cancel' subcommand
	cancelCommand := flag.NewFlagSet("cancel", flag.ExitOnError)

	// Set up 'issue' subcommand
	issueCommand := flag.NewFlagSet("issue", flag.ExitOnError)

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
		stackBuild(*buildBranch, *buildTargets)
	case cancelCommand.Name():
		if errParse := cancelCommand.Parse(os.Args[2:]); errParse != nil {
			log.Fatalf("error parsing cancel flags: %v", errParse)
		}
	case issueCommand.Name():
		if errParse := issueCommand.Parse(os.Args[2:]); errParse != nil {
			log.Fatalf("error parsing issue flags: %v", errParse)
		}
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

	// Execute on 'issue' subcommand
	if issueCommand.Parsed() {
		fmt.Println("'issue' is not yet implemented.") // TODO
		os.Exit(0)

		if issueCommand.NFlag() == 0 {
			fmt.Println("Please specify an issue to create/update.")
			issueCommand.PrintDefaults()
			os.Exit(1)
		}

		issueCommand.Visit(func(f *flag.Flag) {})
	}
}
