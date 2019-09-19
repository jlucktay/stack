package util

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// CurrentBranch parses out the current git branch.
func CurrentBranch() (s string) {
	gitRaw, errExec := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if errExec == nil {
		s = strings.TrimSpace(string(gitRaw))
	}
	return
}

func MustHaveZeroUnpushedCommits() {
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
		fmt.Printf("You have %d unpushed commit(s) on the '%s' branch!\n%v", nUnpushed, CurrentBranch(), yeahNah)
		os.Exit(1)
	}
}

// By special request from one Mr Richard Weston.
const yeahNah = `
! _____.___.             .__       _______         .__
! \__  |   | ____ _____  |  |__    \      \ _____  |  |__
!  /   |   |/ __ \\__  \ |  |  \   /   |   \\__  \ |  |  \
!  \____   \  ___/ / __ \|   Y  \ /    |    \/ __ \|   Y  \
!  / ______|\___  >____  /___|  / \____|__  (____  /___|  /
!  \/           \/     \/     \/          \/     \/     \/
`
