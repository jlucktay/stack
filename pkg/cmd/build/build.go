package build

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jlucktay/stack/pkg/common"
	"github.com/jlucktay/stack/pkg/internal/util"
)

// NewCommand returns the build command.
func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "build",
		Short: "Queue a plan to build this Terraform stack",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("'%s' called, with extraneous args: %s\n", cmd.CalledAs(), args)

			branch, errBranch := cmd.Flags().GetString("branch")
			if errBranch != nil {
				panic(errBranch)
			}

			target, errTarget := cmd.Flags().GetString("target")
			if errTarget != nil {
				panic(errTarget)
			}

			common.StackQueue(
				branch,
				target,
				viper.GetUint("azureDevOps.buildDefID"),
			)
		},
	}

	c.Flags().StringP("branch", "b", util.CurrentBranch(), "If given, plan from this git branch.\n"+
		"Defaults to the current branch.")
	c.Flags().StringP("target", "t", "", "If given, target these specific Terraform resources only.\n"+
		"Delimit multiple target IDs with a semi-colon ';'.")

	return c
}
