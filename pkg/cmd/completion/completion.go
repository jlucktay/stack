package completion

import (
	"os"

	"github.com/spf13/cobra"
)

// NewCommand returns the completion command.
func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "completion [bash|zsh|powershell]",
		Short: "Generate completion script",
		Long: `To load completions:

Bash:

$ source <(stack completion bash)

# To load completions for each session, execute once:
Linux:
	$ stack completion bash > /etc/bash_completion.d/stack
MacOS:
	$ stack completion bash > /usr/local/etc/bash_completion.d/stack

Zsh:

$ source <(stack completion zsh)

# To load completions for each session, execute once:
	$ stack completion zsh > "${fpath[1]}/_stack"
`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "powershell"},
		Args:                  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var errGen error

			switch args[0] {
			case "bash":
				errGen = cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				errGen = cmd.Root().GenZshCompletion(os.Stdout)
			case "powershell":
				errGen = cmd.Root().GenPowerShellCompletion(os.Stdout)
			}

			if errGen != nil {
				panic(errGen)
			}
		},
	}
}
