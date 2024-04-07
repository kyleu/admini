// Content managed by Project Forge, see [projectforge.md] for details.
package main // import admini.dev/admini

import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/cmd"
)

var (
	version = "0.4.2" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	cmd.Entrypoint(&app.BuildInfo{Version: version, Commit: commit, Date: date})
}
