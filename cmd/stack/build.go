package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const yeahNah = `
! _____.___.             .__       _______         .__
! \__  |   | ____ _____  |  |__    \      \ _____  |  |__
!  /   |   |/ __ \\__  \ |  |  \   /   |   \\__  \ |  |  \
!  \____   \  ___/ / __ \|   Y  \ /    |    \/ __ \|   Y  \
!  / ______|\___  >____  /___|  / \____|__  (____  /___|  /
!  \/           \/     \/     \/          \/     \/     \/`

func stackBuild(branch string) {
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
	common.getPAT()
	pat, errPat := common.getPAT()
	common.getPAT()
	common.getPAT()
	if errPat != nil {
		log.Fatal(errPat)
	}

	// 2
	stackPath, errStackPath := getStackPath()
	if errStackPath != nil {
		log.Fatal(errStackPath)
	}

	// 3
	payload, errPayload := getPostPayload(5, stackPath, *buildTargets, *buildBranch)
	if errPayload != nil {
		log.Fatal(errPayload)
	}

	// 4
	apiResult, errAPI := sendBuildRequest(pat, payload)
	if errAPI != nil {
		log.Fatal(errAPI)
	}

	// 5
	fmt.Println("Build URL:", apiResult)
}
