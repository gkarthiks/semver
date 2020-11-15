package utils

import "github.com/Masterminds/semver/v3"

var (
	// DefaultVersion holds the default version value for error returns
	DefaultVersion, _ = semver.NewVersion("0.0.0")
	// Verbose defines the global variable for the verbose flag
	Verbose      bool
	BuildVersion = "development"
)
