package destroy

import (
	"github.com/spf13/cobra"

	"go.jlucktay.dev/stack/pkg/cmd"
)

// NewCommand returns the destroy command.
func NewCommand() *cobra.Command {
	return cmd.NewQueueCommand("destroy")
}
