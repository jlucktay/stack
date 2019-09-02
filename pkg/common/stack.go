package common

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GetStackPath will split the current working directory on 'prefix' and then check if it is part of a git repository
// with 'remote' set as a remote.
func GetStackPath(prefix, remote string) (string, error) {
	wd, errWd := os.Getwd()
	if errWd != nil {
		return "", errWd
	}

	xwd := strings.Split(wd, prefix)
	if len(xwd) < 2 {
		return "", fmt.Errorf("current working directory '%s' is not under '%s'", wd, prefix)
	}

	errRemotes := validateGitRemotes(remote)
	if errRemotes != nil {
		return "", errRemotes
	}

	return xwd[1], nil
}

// validateGitRemotes takes a string argument and searches for it in all remotes that the current git repository
// fetches from. If the string is not found, an error is returned, otherwise nil.
func validateGitRemotes(needle string) error {
	remotes, errExec := exec.Command("git", "remote", "--verbose").Output()
	if errExec != nil {
		return errExec
	}

	found := false

	for _, remote := range strings.Split(string(remotes), "\n") {
		if strings.Contains(remote, "(fetch)") &&
			(strings.Contains(remote, needle) || strings.Contains(strings.ReplaceAll(remote, ":", "/"), needle)) {
			found = true
			break
		}
	}

	if found {
		return nil
	}

	return fmt.Errorf("current git repo does not fetch from '%s' as a remote", needle)
}
