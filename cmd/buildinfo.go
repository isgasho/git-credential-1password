package cmd // nolint:golint // see cmd.go

import "fmt"

const (
	// DefaultVersion is the default 'DEV' version string
	DefaultVersion = "~dev~"
)

var (
	// Version is the application's version
	Version string // nolint:gochecknoglobals // needs to be public
	// Build is the current git Commit
	Build string // nolint:gochecknoglobals // needs to be public
)

func getVersion() string {
	if Version == "" {
		// initialize version with default 'DEV' version string
		// this init() is required because a CI/CD build might
		// not always set a proper non-empty string
		Version = DefaultVersion
	}

	if Build != "" {
		return fmt.Sprintf("%s (build #%s)", Version, Build)
	}

	return Version
}
