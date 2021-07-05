package cmd

import (
	"github.com/kyleu/admini/app/controller"
	"go.uber.org/zap"
)

func startSite(flags *Flags, logger *zap.SugaredLogger) (*zap.SugaredLogger, error) {
	r := controller.SiteRoutes()
	// f := filesystem.NewFileSystem(flags.ConfigDir, logger)

	controller.SetState(nil, logger)

	return webserver(flags.Addr(), logger, r)
}

