package init

import (
	"github.com/spf13/cobra"

	"go.jlucktay.dev/stack/pkg/common"
)

// NewCommand returns the init command.
func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "init",
		Short: "Initialise this Terraform stack against remote state",
		Long: `Runs Terraform with values derived from configured settings to initialise the
current stack directory, using the Azure storage account for the remote state
backend.`,
		Run: func(_ *cobra.Command, _ []string) {
			common.InitStack()
		},
	}

	return c
}
