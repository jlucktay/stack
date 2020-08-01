// Package cmd contains the root of the CLI command section of our migration support tool, 'stack', which leverages
// logic from other packages stored elsewhere in the repo.
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"go.jlucktay.dev/stack/internal/exit"
	"go.jlucktay.dev/stack/pkg/cmd/build"
	"go.jlucktay.dev/stack/pkg/cmd/cancel"
	"go.jlucktay.dev/stack/pkg/cmd/completion"
	"go.jlucktay.dev/stack/pkg/cmd/destroy"
	stackinit "go.jlucktay.dev/stack/pkg/cmd/init"
	"go.jlucktay.dev/stack/pkg/cmd/issue"
	"go.jlucktay.dev/stack/pkg/cmd/version"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "stack",
		Short: "A support tool for working with Terraform stacks, Azure DevOps pipelines, and GitHub projects/repos.",
		Long: `A support tool for working with Terraform stacks, Azure DevOps pipelines, and GitHub projects/repos.

Stack was built to enable quicker turnaround time while working with Terraform stacks that were built via Azure DevOps
pipelines, primarily to avoid the sluggish and generally awful UI of the latter.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		//	Run: func(cmd *cobra.Command, args []string) { },
	}

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(
		build.NewCommand(),
		cancel.NewCommand(),
		completion.NewCommand(),
		destroy.NewCommand(),
		stackinit.NewCommand(),
		issue.NewCommand(),
		version.NewCommand(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(exit.RootExecute)
	}
}

const (
	viperConfigName = "stack.config"
	viperConfigType = "json"
	viperConfigFile = viperConfigName + "." + viperConfigType
)

// initConfig reads in a config file from a preference order of search paths, and also environment variables if set.
func initConfig() {
	// Find home directory.
	home, errHome := homedir.Dir()
	if errHome != nil {
		fmt.Println(errHome)
		os.Exit(exit.HomeNotFound)
	}

	if strings.Contains(home, "/tmp/") {
		return
	}

	// Name of config file, and extension.
	viper.SetConfigName(viperConfigName)
	viper.SetConfigType(viperConfigType)

	// Paths to search (in descending order of preference) for the config file.
	configPaths := []string{
		".",
		filepath.Join(home, ".config/stack"),
		"/etc/stack",
	}

	for _, configPath := range configPaths {
		viper.AddConfigPath(configPath)
	}

	// Read in environment variables that match.
	viper.AutomaticEnv()

	if errViperRead := viper.ReadInConfig(); errViperRead != nil {
		if _, ok := errViperRead.(viper.ConfigFileNotFoundError); ok {
			return
		}

		panic(fmt.Sprintf("Fatal error reading config file '%s':\n%s\n", viper.ConfigFileUsed(), errViperRead))
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
