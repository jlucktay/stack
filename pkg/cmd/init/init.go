package init

import (
	"github.com/spf13/cobra"

	"github.com/jlucktay/stack/pkg/common"
)

// NewCommand returns the init command.
func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "init",
		Short: "Initialise this Terraform stack against remote state",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			common.InitStack()
		},
	}

	return c
}
