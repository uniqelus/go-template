package pkgversion

import (
	"fmt"
	"runtime"
)

// These variables are set at build time using ldflags.
var (
	Version   = "dev"
	GitCommit = "none"
	BuildTime = "unknown"
	GoVersion = runtime.Version()
	OS        = runtime.GOOS
	Arch      = runtime.GOARCH
)

// Info returns a struct with all version information.
type Info struct {
	Version   string `json:"version"`
	GitCommit string `json:"git_commit"`
	BuildTime string `json:"build_time"`
	GoVersion string `json:"go_version"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
}

// Get returns the current version info.
func Get() Info {
	return Info{
		Version:   Version,
		GitCommit: GitCommit,
		BuildTime: BuildTime,
		GoVersion: GoVersion,
		OS:        OS,
		Arch:      Arch,
	}
}

// String returns a formatted version string.
func String() string {
	return fmt.Sprintf(
		"Version: %s\nGit Commit: %s\nBuild Time: %s\nGo Version: %s\nOS/Arch: %s/%s",
		Version, GitCommit, BuildTime, GoVersion, OS, Arch,
	)
}
