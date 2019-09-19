package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jlucktay/stack/pkg/version"
)

// NewCommand returns the version command.
func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "version",
		Short: "Show details of this binary's current version",
		Long: `Show details of this binary's current version.

The version value shown follows semantic versioning: https://semver.org
Commit is the SHA1 hash of the git commit built from.
Date is the timestamp of the build.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.Details())
		},
	}

	return c
}
