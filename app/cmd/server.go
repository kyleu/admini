// Package cmd: $PF_IGNORE$
package cmd

import (
	"fmt"
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
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const keyServer = "server"

func serverCmd() *cobra.Command {
	short := fmt.Sprintf("Starts the http server on port %d (by default)", util.AppPort)
	f := func(*cobra.Command, []string) error { return startServer() }
	ret := &cobra.Command{Use: keyServer, Short: short, RunE: f}
	return ret
}

func startServer() error {
	err := initIfNeeded()
	if err != nil {
		return errors.Wrap(err, "error initializing application")
	}

	r, _, err := loadServer(_flags, _logger)
	if err != nil {
		return err
	}

	_, err = listenandserve(util.AppName, _flags.Address, _flags.Port, r)
	return err
}

func loadServer(flags *Flags, logger *zap.SugaredLogger) (*router.Router, *zap.SugaredLogger, error) {
	r := controller.AppRoutes()

	f := filesystem.NewFileSystem(flags.ConfigDir, logger)
	ls := loader.NewService()
	ls.Set(schema.OriginPostgres, lpostgres.NewLoader(logger))
	if database.SQLiteEnabled {
		ls.Set(schema.OriginSQLite, lsqlite.NewLoader(logger))
	}
	ls.Set(schema.OriginMock, lmock.NewLoader(logger))

	st, err := app.NewState(flags.Debug, _buildInfo, r, f, logger)
	if err != nil {
		return nil, logger, err
	}

	ss := source.NewService(f, ls, logger)
	st.Sources = ss
	st.Projects = project.NewService(f, ss, ls, logger)
	st.Loaders = ls

	controller.SetAppState(st, logger)

	logger.Infof("started %s using address [%s:%d] on %s:%s", util.AppName, flags.Address, flags.Port, runtime.GOOS, runtime.GOARCH)

	return r, logger, nil
}
