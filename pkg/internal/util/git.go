package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// CurrentBranch parses out the name of the current git branch, if we are inside a git repo.
// Otherwise, an empty string is returned.
func CurrentBranch() string {
	cmdGit := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	_, errRepoTest := cmdGit.Output()

	// With thanks to:
	// https://stackoverflow.com/questions/10385551/get-exit-code-go
	if errRepoTest != nil {
		if errExit, ok := errRepoTest.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0
			if status, ok := errExit.Sys().(syscall.WaitStatus); ok {
				exitStatus := status.ExitStatus()
				if exitStatus == 128 {
					return ""
				}

				log.Fatalf("'%+v' exit code: %d", cmdGit, exitStatus)
			}
		}

		log.Fatalf("'%+v': %+v", cmdGit, errRepoTest)
	}

	gitRaw, errBranch := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if errBranch != nil {
		panic(errBranch)
	}

	return strings.TrimSpace(string(gitRaw))
}

func MustHaveZeroUnpushedCommits(targetBranch string) {
	local := mustGetCommitHash(targetBranch)
	remote := mustGetCommitHash("origin/" + targetBranch)
	commitRange := fmt.Sprintf("%s...%s", remote, local)
	rawCommitCount, errLog := exec.Command("git", "log", "--pretty=oneline", commitRange).CombinedOutput()

	if errLog != nil {
		fmt.Printf("Error counting commits between %s and %s commits:\n%s\n", remote, local, rawCommitCount)

		if strings.Contains(errLog.Error(), "exit status 128") {
			fmt.Printf("error counting unpushed commits; check to confirm that %s exists on the remote\n", targetBranch)
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
