package destroy

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewCommand returns the destroy command.
func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "destroy",
		Short: "Queue a plan to destroy this Terraform stack",
		Long: `A longer description that spans multiple lines and likely contains examples
		and usage of using your command. For example:

		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("destroy called")
		},
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// destroyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// destroyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return c
}
