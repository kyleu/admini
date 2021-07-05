package cmd

import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller"
	"github.com/kyleu/admini/app/filesystem"
	"go.uber.org/zap"
)

func startSite(flags *Flags, logger *zap.SugaredLogger) (*zap.SugaredLogger, error) {
	r := controller.SiteRoutes()
	f := filesystem.NewFileSystem(flags.ConfigDir, logger)
	st, err := app.NewState(flags.Debug, AppBuildInfo, r, f, nil, logger)
	if err != nil {
		return logger, err
	}
	controller.SetState(st, logger)
	return webserver(flags.Addr(), logger, r)
}

