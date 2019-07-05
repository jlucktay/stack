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

func validateGitRemotes(needle string) error {
	remotes, errExec := exec.Command("git", "remote", "--verbose").Output()
	if errExec != nil {
		return errExec
	}

	for _, remote := range strings.Split(string(remotes), "\n") {
		if strings.Contains(remote, "(fetch)") && !strings.Contains(remote, needle) {
			return fmt.Errorf("current git repo does not fetch from '%s' as a remote", needle)
		}
	}

	return nil
}
