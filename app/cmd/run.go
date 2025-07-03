package cmd

import (
	"context"

	"admini.dev/admini/app"
	"admini.dev/admini/app/util"
)

func Run(ctx context.Context, bi *app.BuildInfo) (util.Logger, error) {
	_buildInfo = bi

	if err := rootCmd(ctx).Execute(); err != nil {
		return util.RootLogger, err
	}
	return util.RootLogger, nil
}
