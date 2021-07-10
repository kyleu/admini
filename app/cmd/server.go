package cmd

import (
	"runtime"

	"github.com/fasthttp/router"
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller"
	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/loader/lmock"
	"github.com/kyleu/admini/app/loader/lpostgres"
	"github.com/kyleu/admini/app/loader/lsqlite"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/app/util"
	"go.uber.org/zap"
)

func startServer(flags *Flags, logger *zap.SugaredLogger) (*zap.SugaredLogger, error) {
	r, logger, err := loadServer(flags, logger)
	if err != nil {
		return nil, err
	}

	_, err = listenandserve(util.AppName, flags.Address, flags.Port, r)
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func loadServer(flags *Flags, logger *zap.SugaredLogger) (*router.Router, *zap.SugaredLogger, error) {
	r := controller.BuildRouter()

	f := filesystem.NewFileSystem(flags.ConfigDir, logger)
	ls := loader.NewService()
	ls.Set(schema.OriginPostgres, lpostgres.NewLoader(logger))
	if database.SQLiteEnabled {
		ls.Set(schema.OriginSQLite, lsqlite.NewLoader(logger))
	}
	ls.Set(schema.OriginMock, lmock.NewLoader(logger))

	st, err := app.NewState(flags.Debug, AppBuildInfo, r, f, logger)
	if err != nil {
		return nil, logger, err
	}

	ss := source.NewService(f, ls, logger)
	st.Sources = ss
	st.Projects = project.NewService(f, ss, ls, logger)
	st.Loaders = ls

	controller.SetAppState(st, logger)

	logger.Infof("started using address [%s:%d] on %s:%s", flags.Address, flags.Port, runtime.GOOS, runtime.GOARCH)

	return r, logger, nil
}
