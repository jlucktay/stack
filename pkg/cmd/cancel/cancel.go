package cancel

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cancelCmd represents the cancel command
func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "cancel",
		Short: "Cancel release(s) of built/planned Terraform stack",
		Long: `A longer description that spans multiple lines and likely contains examples
		and usage of using your command. For example:

		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("'%s' called; not yet implemented!\n", cmd.CalledAs())
		},
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cancelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return c
}

// cancel flow:
// 0. parse optional arguments
// 1. get PAT from XDG dir
// 2. build POST payload
// 3. get all my active releases
// 4. for each release:
// 4.1 cancel
// 4.2 abandon
// 5. print IDs of cancelled/abandoned releases
