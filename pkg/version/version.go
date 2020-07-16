package version

import (
	"fmt"
)

// Take ldflags from GoReleaser.
var (
	version, commit, goVersion, date, builtBy string //nolint:gochecknoglobals
)

func Details() string {
	return fmt.Sprintf("stack %s built from commit %s with %s on %s by %s.", version, commit, goVersion, date, builtBy)
}
