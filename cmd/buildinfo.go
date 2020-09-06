package cmd

import (
	"fmt"
)

var (
	// Version is the application's version.
	Version string // nolint:gochecknoglobals // needs to be public
	// Build is the current git Commit.
	Build string // nolint:gochecknoglobals // needs to be public
	// Date is the current date.
	Date string // nolint:gochecknoglobals // needs to be public
)

func getVersion() string {
	if Version == "" {
		// initialize version with default 'DEV' version string
		// this init() is required because a CI/CD build might
		// not always set a proper non-empty string
		Version = "~dev~"
	}

	if Build != "" {
		return fmt.Sprintf("%s (build #%s, %s)", Version, Build, Date)
	}

	return Version
}
