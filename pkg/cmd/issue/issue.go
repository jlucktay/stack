package issue

import (
	"github.com/spf13/cobra"

	"github.com/jlucktay/stack/pkg/common"
)

// NewCommand returns the issue command.
func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "issue",
		Short: "Add a GitHub issue for this Terraform stack",
		Long: `Creates an issue in the configured GitHub org/repo using the provided PAT, by
way of an interactive text editor from the command line.

The 'title' flag for this subcommand is optional.

Example usage:
$ stack issue --title "My issue title"

The above command would create a new issue in the configured GitHub org/repo
titled "My issue title" with body text entered and saved through the user's
default editor, denoted by the EDITOR environment variable.`,

		Run: func(cmd *cobra.Command, args []string) {
			title, errTitle := cmd.Flags().GetString("title")
			if errTitle != nil {
				panic(errTitle)
			}
			if title == "" {
				title = "New issue"
			}

			common.CreateIssue(title)
		},
	}

	c.Flags().StringP("title", "t", "", "If given, title the issue with this string.")

	return c
}
