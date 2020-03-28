package common

import (
	"fmt"
	"log"
	"strings"

	"github.com/jlucktay/stack/pkg/internal/util"

	"github.com/spf13/viper"
)

// queue flow:
// 0.1. count unpushed commits and warn if > 0
// 1. get stack path
// 2. build POST payload using parameters
// 3. send request to API
// 4. print URL of build from API result

const (
	azureTargetDelimiter = ";"
	targetArgDelimiter   = ","
)

//nolint:funlen // TODO
func StackQueue(branch, targets string, defID uint) {
	buildQueueing := fmt.Sprintf("Queueing build def %d from branch '%s' ", defID, branch)

	if len(targets) == 0 {
		buildQueueing += "for all available Terraform targets.\n"
	} else {
		buildQueueing += "scoped to the following Terraform target(s):\n"

		for _, targ := range strings.Split(targets, targetArgDelimiter) {
			buildQueueing += fmt.Sprintf(" - %s\n", targ)
		}
	}

	log.Println(buildQueueing)

	// 0.1
	util.MustHaveZeroUnpushedCommits(branch)

	// 1
	stackPath := mustGetStackPath()

	// 2
	parameters := make(map[string]string)
	parameters["StackPath"] = stackPath

	if len(targets) > 0 {
		xTargets := strings.Split(targets, targetArgDelimiter)
		parameters["TerraformTarget"] = strings.Join(xTargets, azureTargetDelimiter)
	}

	payload, errPayload := GetPostPayload(defID, parameters, branch)
	if errPayload != nil {
		panic(errPayload)
	}

	// 3
	adoOrgKey := "azureDevOps.org"
	if !viper.IsSet(adoOrgKey) {
		panic("the Azure DevOps organisation has not been specified under '" + adoOrgKey + "' in your config")
	}

	adoProjectKey := "azureDevOps.project"
	if !viper.IsSet(adoProjectKey) {
		panic("the Azure Devops project has not been specified under '" + adoProjectKey + "' in your config")
	}

	adoPATKey := "azureDevOps.pat"
	if !viper.IsSet(adoPATKey) {
		panic("the Azure DevOps personal access token has not been specified under '" + adoPATKey + "' in your config")
	}

	apiURL := fmt.Sprintf(
		"https://dev.azure.com/%s/%s/_apis/build/builds?api-version=5.0",
		viper.GetString(adoOrgKey),
		viper.GetString(adoProjectKey),
	)

	req, errReq := CreateBuildRequest(apiURL, viper.GetString(adoPATKey), payload)
	if errReq != nil {
		panic(errReq)
	}

	apiResult, errAPI := SendBuildRequest(req)
	if errAPI != nil {
		panic(errAPI)
	}

	// 4
	fmt.Println("Stack (plan) URL:", apiResult)
}
