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
	return fmt.Sprintf("stack %s built from commit %s with %s on %s by %s.",
		version, commit, runtime.Version(), date, builtBy)
}
