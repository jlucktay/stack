package version

import (
	"fmt"
	"runtime"
)

// Take ldflags from GoReleaser.
var (
	version, commit, date, builtBy string //nolint:gochecknoglobals
)

func Details() string {
	output := "stack"

	if len(version) > 0 {
		output += fmt.Sprintf(" %s", version)
	}

	output += " built"

	if len(commit) > 0 {
		output += fmt.Sprintf(" from commit %s", commit)
	}

	output += fmt.Sprintf(" with %s", runtime.Version())

	if len(date) > 0 {
		output += fmt.Sprintf(" on %s", date)
	}

	if len(builtBy) > 0 {
		output += fmt.Sprintf(" by %s", builtBy)
	}

	output += "."

	return output
}
