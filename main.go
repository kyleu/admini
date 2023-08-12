// Content managed by Project Forge, see [projectforge.md] for details.
package main // import admini.dev/admini

import (
	"os"

	"admini.dev/admini/app"
	"admini.dev/admini/app/cmd"
	"admini.dev/admini/app/lib/log"
)

var (
	version = "0.2.25" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	logger, err := cmd.Run(&app.BuildInfo{Version: version, Commit: commit, Date: date})
	if err != nil {
		const msg = "exiting due to error"
		if logger == nil {
			println(log.Red.Add(err.Error())) //nolint:forbidigo
			println(log.Red.Add(msg))         //nolint:forbidigo
		} else {
			logger.Error(err)
			logger.Error(msg)
		}
		os.Exit(1)
	}
}
