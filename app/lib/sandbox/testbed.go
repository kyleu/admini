// Package sandbox $PF_IGNORE$
package sandbox

import (
	"context"

	"go.uber.org/zap"

	"admini.dev/admini/app"
	"admini.dev/admini/app/util"
)

var testbed = &Sandbox{Key: "testbed", Title: "Testbed", Icon: "star", Run: onTestbed}

func onTestbed(ctx context.Context, st *app.State, logger *zap.SugaredLogger) (any, error) {
	ret := util.ValueMap{"status": "ok"}
	return ret, nil
}
