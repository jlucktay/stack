package version

import (
	"fmt"
)

// Take ldflags from GoReleaser.
var (
	Version, Commit, GoVersion, Date, BuiltBy string //nolint:gochecknoglobals
)

func Details() string {
	return fmt.Sprintf("stack %s built from commit %s with %s on %s by %s.", Version, Commit, GoVersion, Date, BuiltBy)
}
