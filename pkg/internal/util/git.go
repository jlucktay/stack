package util

import (
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
