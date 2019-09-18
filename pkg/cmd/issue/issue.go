package issue

import (
	"fmt"

	"github.com/spf13/cobra"
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
			fmt.Println("issue called")
		},
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// issueCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// issueCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return c
}
