package issue

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jlucktay/stack/pkg/common"
)

// NewCommand returns the issue command.
func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "issue",
		Short: "Add a GitHub issue for this Terraform stack",
		Long: `Creates an issue in the configured GitHub org/repo using the provided PAT.

The 'title' flag for this subcommand is optional.

Example usage:
$ stack issue --title "My issue title" I found a problem with this stack

The above command would create a new issue in the configured GitHub org/repo
titled "My issue title" with body text of "I found a problem with this stack".`,
		Run: func(cmd *cobra.Command, args []string) {
			title, errTitle := cmd.Flags().GetString("title")
			if errTitle != nil {
				panic(errTitle)
			}
			if title == "" {
				title = "New issue"
			}

			if len(args) == 0 {
				fmt.Println("No issue text was given!")
				cmd.UsageString()
				os.Exit(1)
			}

			common.CreateIssue(title, args...)
		},
	}

	c.Flags().StringP("title", "t", "", "If given, title the issue with this string.")

	return c
}
