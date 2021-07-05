package cmd

import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller"
	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/loader/lmock"
	"github.com/kyleu/admini/app/loader/lpostgres"
	"github.com/kyleu/admini/app/loader/lsqlite"
	"github.com/kyleu/admini/app/schema"
	"go.uber.org/zap"
)

func startApp(flags *Flags, logger *zap.SugaredLogger) (*zap.SugaredLogger, error) {
	r := controller.BuildRouter()

	f := filesystem.NewFileSystem(flags.ConfigDir, logger)
	ls := loader.NewService()
	ls.Set(schema.OriginPostgres, lpostgres.NewLoader(logger))
	if database.SQLiteEnabled {
		ls.Set(schema.OriginSQLite, lsqlite.NewLoader(logger))
	}
	ls.Set(schema.OriginMock, lmock.NewLoader(logger))

	st, err := app.NewState(flags.Debug, AppBuildInfo, r, f, ls, logger)
	if err != nil {
		return logger, err
	}
	controller.SetState(st, logger)

	return webserver(flags.Addr(), logger, r)
}

