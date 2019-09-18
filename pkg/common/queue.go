package common

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

const yeahNah = `
! _____.___.             .__       _______         .__
! \__  |   | ____ _____  |  |__    \      \ _____  |  |__
!  /   |   |/ __ \\__  \ |  |  \   /   |   \\__  \ |  |  \
!  \____   \  ___/ / __ \|   Y  \ /    |    \/ __ \|   Y  \
!  / ______|\___  >____  /___|  / \____|__  (____  /___|  /
!  \/           \/     \/     \/          \/     \/     \/`

// queue flow:
// 0.1. count unpushed commits and warn if > 0
// 1. get stack path
// 2. build POST payload using parameters
// 3. send request to API
// 4. print URL of build from API result

//nolint
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
	unpushedRaw, errExec := exec.Command("git", "rev-list", "--count", "@{u}..").Output()
	if errExec != nil {
		if errExec.Error() == "exit status 128" {
			fmt.Println("error counting unpushed commits; check to confirm there is an upstream configured")
		}
		panic(errExec)
	}

	nUnpushed, errAtoi := strconv.Atoi(strings.TrimSpace(string(unpushedRaw)))
	if errAtoi != nil {
		panic(errAtoi)
	}
	if nUnpushed > 0 {
		panic(fmt.Sprintf("you have %d unpushed commit(s)) on the '%s' branch:\n%v", nUnpushed, branch, yeahNah))
	}

	// 1
	stackPath := MustGetStackPath()

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
