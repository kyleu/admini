package cmd

import (
	"runtime"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/util"
	"go.uber.org/zap"
)

func startSite(flags *Flags, logger *zap.SugaredLogger) (*zap.SugaredLogger, error) {
	r := controller.SiteRoutes()
	f := filesystem.NewFileSystem(flags.ConfigDir, logger)
	st, err := app.NewState(flags.Debug, AppBuildInfo, r, f, logger)
	if err != nil {
		return logger, err
	}
	controller.SetSiteState(st, logger)
	logger.Infof("marketing site started using address [%s] on %s:%s", flags.Addr(), runtime.GOOS, runtime.GOARCH)
	_, err = listenandserve(util.AppName+"-site", flags.Address, flags.Port, r)
	if err != nil {
		return nil, err
	}
	return logger, nil
}
