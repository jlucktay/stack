package util

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"

	"go.jlucktay.dev/stack/internal/exit"
)

func mustGetRepo() *git.Repository {
	dir, errGetWD := os.Getwd()
	if errGetWD != nil {
		panic(errGetWD)
	}

	repository, errOpen := git.PlainOpen(dir)
	if errOpen != nil {
		if errors.Is(errOpen, git.ErrRepositoryNotExists) {
			return nil
		}

		panic(errOpen)
	}

	return repository
}

// CurrentBranch parses out the name of the current git branch, if we are inside a git repo.
// Otherwise, an empty string is returned.
func CurrentBranch() string {
	ref, errRef := mustGetRepo().Reference(plumbing.HEAD, true)
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

	if !strings.EqualFold(local, remote) {
		fmt.Printf("Local vs remote branch '%s' has unpushed commits!\n%v", targetBranch, yeahNah)
		os.Exit(exit.UnpushedCommits)
	}
}

func mustGetCommitHash(branch string) string {
	ref, errRef := mustGetRepo().Reference(plumbing.NewBranchReferenceName(branch), true)
	if errRef != nil {
		panic(errRef)
	}

	return ref.Hash().String()
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
