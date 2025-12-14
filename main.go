package main // import admini.dev/admini

import (
	"context"

	"admini.dev/admini/app"
	"admini.dev/admini/app/cmd"
)

var (
	version = "0.6.5" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	cmd.Entrypoint(context.Background(), &app.BuildInfo{Version: version, Commit: commit, Date: date})
}
