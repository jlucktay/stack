package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jlucktay/stack/pkg/common"
	"github.com/jlucktay/stack/pkg/internal/util"
)

func NewQueueCommand(defType string) *cobra.Command {
	c := &cobra.Command{
		Use:   defType,
		Short: fmt.Sprintf("Queue a plan to %s this Terraform stack", defType),
		Long: fmt.Sprintf(`This command queues a build on Azure DevOps to %s this Terraform stack.

Configured values for the Azure DevOps project, organisation, PAT, and build
definition ID are all used, as well as the stack prefix value to compose the
key of the Terraform state file within the Azure storage account.

Example usage:
$ stack %[1]s --branch feature/new-stack --target "azurerm_virtual_machine.example,azurerm_resource_group.example"
Stack (plan) URL: https://dev.azure.com/Org/Project/_build/results?buildId=1234`, defType),
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
				viper.GetUint(fmt.Sprintf("azureDevOps.%sDefID", defType)),
			)
		},
	}

	c.Flags().StringP("branch", "b", util.CurrentBranch(), "If given, plan from this git branch.\n"+
		"Defaults to the current branch.")
	c.Flags().StringP("target", "t", "", "If given, target these specific Terraform resources only.\n"+
		"Delimit multiple target IDs with a comma ','.")

	return c
}
