package version

import (
	"fmt"
)

// Take ldflags from GoReleaser
var (
	//nolint
	version, commit, date, builtby string
)

func Details() string {
	return fmt.Sprintf("stack v%s from commit %s, built %s by %s.", version, commit, date, builtby)
}