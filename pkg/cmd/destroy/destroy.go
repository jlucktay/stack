package destroy

import (
	"github.com/spf13/cobra"

	"github.com/jlucktay/stack/pkg/cmd"
)

// NewCommand returns the destroy command.
func NewCommand() *cobra.Command {
	return cmd.NewQueueCommand("destroy")
}
