package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/jlucktay/stack/pkg/common"
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

func stackQueue(branch, targets string, defID uint) {
	// 0.1
	unpushedRaw, errExec := exec.Command("git", "rev-list", "--count", "@{u}..").Output()
	if errExec != nil {
		log.Fatal(errExec)
	}

	nUnpushed, errAtoi := strconv.Atoi(strings.TrimSpace(string(unpushedRaw)))
	if errAtoi != nil {
		log.Fatal(errAtoi)
	}
	if nUnpushed > 0 {
		log.Fatalf("you have %d unpushed commit(s) on the '%s' branch:\n%v", nUnpushed, branch, yeahNah)
	}

	// 1
	stackPath, errStackPath := common.GetStackPath(
		viper.GetString("stackPrefix"),
		fmt.Sprintf(
			"github.com/%s/%s",
			viper.GetString("github.org"),
			viper.GetString("github.repo"),
		),
	)
	if errStackPath != nil {
		log.Fatal(errStackPath)
	}

	// 2
	parameters := make(map[string]string)
	parameters["StackPath"] = stackPath

	if len(targets) > 0 {
		parameters["TerraformTarget"] = targets
	}

	payload, errPayload := common.GetPostPayload(defID, parameters, branch)
	if errPayload != nil {
		log.Fatal(errPayload)
	}

	// 3
	apiURL := fmt.Sprintf(
		"https://dev.azure.com/%s/%s/_apis/build/builds?api-version=5.0",
		viper.GetString("azureDevOps.org"),
		viper.GetString("azureDevOps.project"),
	)
	apiResult, errAPI := common.SendBuildRequest(
		apiURL,
		viper.GetString("azureDevOps.pat"),
		payload,
	)
	if errAPI != nil {
		log.Fatal(errAPI)
	}

	// 4
	fmt.Println("Stack (plan) URL:", apiResult)
}
