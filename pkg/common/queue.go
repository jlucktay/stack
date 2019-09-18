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

func StackQueue(branch, targets string, defID uint) {
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
		log.Fatalf("you have %d unpushed commit(s) on the '%s' branch:\n%v", nUnpushed, branch, yeahNah)
	}

	// 1
	stackPath, errStackPath := GetStackPath(
		viper.GetString("stackPrefix"),
		fmt.Sprintf(
			"github.com/%s/%s",
			viper.GetString("github.org"),
			viper.GetString("github.repo"),
		),
	)
	if errStackPath != nil {
		panic(errStackPath)
	}

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
