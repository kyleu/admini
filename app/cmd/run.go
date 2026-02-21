package cmd

import (
	"context"

	"github.com/charmbracelet/fang"

	"admini.dev/admini/app"
	"admini.dev/admini/app/util"
)

func Run(ctx context.Context, bi *app.BuildInfo) (util.Logger, error) {
	_buildInfo = bi

	if err := fang.Execute(ctx, rootCmd(ctx), fang.WithVersion(bi.Version), fang.WithCommit(bi.Commit)); err != nil {
		return util.RootLogger, err
	}
	return util.RootLogger, nil
}
