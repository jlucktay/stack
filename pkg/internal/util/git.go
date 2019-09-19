package util

import (
	"fmt"
	"os"
	"os/exec"
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

func MustHaveZeroUnpushedCommits(targetBranch string) {
	local := mustGetCommitHash(targetBranch)
	remote := mustGetCommitHash("origin/" + targetBranch)

	rawCommitCount, errLog := exec.Command(
		"git", "log", "--pretty=oneline",
		fmt.Sprintf("%s...%s", remote, local),
	).CombinedOutput()

	if errLog != nil {
		fmt.Printf("Error counting commits between %s and %s commits:\n%s\n", remote, local, rawCommitCount)

		if strings.Contains(errLog.Error(), "exit status 128") {
			fmt.Printf("error counting unpushed commits;"+
				"check to confirm that %s exists on the remote\n",
				targetBranch,
			)
		}

		panic(errLog)
	}

	lineCount := strings.Count(string(rawCommitCount), "\n")

	if lineCount > 0 {
		fmt.Printf("You have %d unpushed commit(s) on the '%s' branch!\n%v", lineCount, targetBranch, yeahNah)
		os.Exit(1)
	}
}

func mustGetCommitHash(branch string) string {
	rawHash, errRevParse := exec.Command("git", "rev-parse", branch).CombinedOutput()
	if errRevParse != nil {
		fmt.Printf("Error parsing branch ref '%s':\n%s\n", branch, rawHash)
		panic(errRevParse)
	}
	return strings.TrimSpace(string(rawHash))
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
