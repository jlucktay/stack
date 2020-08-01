package util

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"

	"go.jlucktay.dev/stack/internal/exit"
)

// CurrentBranch parses out the name of the current git branch, if we are inside a git repo.
// Otherwise, an empty string is returned.
func CurrentBranch() string {
	dir, errGetWD := os.Getwd()
	if errGetWD != nil {
		panic(errGetWD)
	}

	repository, errOpen := git.PlainOpen(dir)
	if errOpen != nil {
		if errors.Is(errOpen, git.ErrRepositoryNotExists) {
			return ""
		}

		panic(errOpen)
	}

	ref, errRef := repository.Reference(plumbing.HEAD, true)
	if errRef != nil {
		panic(errRef)
	}

	if ref.Name().IsBranch() {
		return ref.Name().Short()
	}

	return ref.String()
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
		os.Exit(exit.UnpushedCommits)
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
