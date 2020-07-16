package build

import (
	"github.com/spf13/cobra"

	"go.jlucktay.dev/stack/pkg/cmd"
)

// NewCommand returns the build command.
func NewCommand() *cobra.Command {
	return cmd.NewQueueCommand("build")
}
