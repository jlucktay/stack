package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"go.jlucktay.dev/version"
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

		RunE: func(cmd *cobra.Command, args []string) error {
			v, err := version.Details()
			if err != nil {
				return fmt.Errorf("error looking up version details: %w", err)
			}

			fmt.Println(v)

			return nil
		},
	}

	return c
}
