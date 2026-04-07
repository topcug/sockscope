package version

// Version is overridden at build time via -ldflags
// "-X github.com/topcug/sockscope/pkg/version.Version=..." by
// GoReleaser. The default "dev" tag is what you get from `go build`
// in a working copy.
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)
