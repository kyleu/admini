// Package sandbox $PF_IGNORE$
package sandbox

import (
	"context"

	"admini.dev/admini/app"
	"admini.dev/admini/app/util"
)

var testbed = &Sandbox{Key: "testbed", Title: "Testbed", Icon: "star", Run: onTestbed}

func onTestbed(ctx context.Context, st *app.State, logger util.Logger) (any, error) {
	ret := util.ValueMap{"status": "ok"}
	return ret, nil
}
