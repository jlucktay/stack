package version

import (
	"fmt"
)

//nolint:gochecknoglobals // Take ldflags from GoReleaser.
var (
	Version   string
	Commit    string
	GoVersion string
	Date      string
	BuiltBy   string
)

func Details() string {
	return fmt.Sprintf("stack %s built from commit %s with %s on %s by %s.", Version, Commit, GoVersion, Date, BuiltBy)
}
