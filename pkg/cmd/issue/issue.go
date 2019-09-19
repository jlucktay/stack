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
		Short: "Add/update a GitHub issue for this Terraform stack",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("No issue text was given!")
				cmd.UsageString()
				os.Exit(1)
			}

			common.CreateIssue(args...)
		},
	}

	return c
}
