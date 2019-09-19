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

func StackQueue(branch, targets string, defID uint) {
	buildQueueing := fmt.Sprintf("Queueing build def %d from branch '%s' ", defID, branch)

	if len(targets) == 0 {
		buildQueueing += "for all available Terraform targets.\n"
	} else {
		buildQueueing += "scoped to the following Terraform target(s):\n"

		for _, targ := range strings.Split(targets, ";") {
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
		parameters["TerraformTarget"] = targets
	}

	payload, errPayload := GetPostPayload(defID, parameters, branch)
	if errPayload != nil {
		panic(errPayload)
	}

	// 3
	apiURL := fmt.Sprintf(
		"https://dev.azure.com/%s/%s/_apis/build/builds?api-version=5.0",
		viper.GetString("azureDevOps.org"),
		viper.GetString("azureDevOps.project"),
	)
	apiResult, errAPI := SendBuildRequest(
		apiURL,
		viper.GetString("azureDevOps.pat"),
		payload,
	)
	if errAPI != nil {
		panic(errAPI)
	}

	// 4
	fmt.Println("Stack (plan) URL:", apiResult)
}
