package cmd

import (
	"fmt"
	"runtime"

	"github.com/kyleu/admini/app/database"
	"github.com/pkg/errors"

	"github.com/kirsle/configdir"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/loader/lmock"
	"github.com/kyleu/admini/app/loader/lpostgres"
	"github.com/kyleu/admini/app/loader/lsqlite"
	"github.com/kyleu/admini/app/log"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

type Flags struct {
	Address   string
	Port      int
	ConfigDir string
	Debug     bool
	JSON      bool
}

var AppBuildInfo *app.BuildInfo

func Run() (*zap.SugaredLogger, error) {
	if AppBuildInfo == nil {
		return nil, errors.New("no build info")
	}

	flags := parseFlags()
	if flags.ConfigDir == "" {
		flags.ConfigDir = configdir.LocalConfig(util.AppName)
		_ = configdir.MakePath(flags.ConfigDir)
	}

	logger, err := log.InitLogging(flags.Debug, flags.JSON)
	if err != nil {
		return logger, err
	}

	addr := fmt.Sprintf("%s:%d", flags.Address, util.AppPort)
	startLog := logger.With(zap.Bool("debug", flags.Debug), zap.String("address", flags.Address), zap.Int("port", flags.Port))
	startLog.Infof("[%s v%s] %s", util.AppName, AppBuildInfo.Version, util.AppURL)
	startLog.Infof("starting using address [%s] on %s:%s", addr, runtime.GOOS, runtime.GOARCH)

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

	s := &fasthttp.Server{Handler: r.Handler, Name: util.AppName, ReadBufferSize: 32768}
	return logger, s.ListenAndServe(addr)
}

func parseFlags() *Flags {
	ret := &Flags{}
	pflag.StringVarP(&ret.Address, "addr", "a", "127.0.0.1", "address to listen on, defaults to [127.0.0.1]")
	pflag.IntVarP(&ret.Port, "port", "p", util.AppPort, fmt.Sprintf("port to listen on, defaults to [%d]", util.AppPort))
	pflag.StringVarP(&ret.ConfigDir, "dir", "d", "", "directory for configuration, defaults to [~/???]")
	pflag.BoolVarP(&ret.Debug, "verbose", "v", false, "enables verbose logging and additional checks")
	pflag.BoolVarP(&ret.JSON, "json", "j", false, "enables json logging")
	pflag.Parse()
	return ret
}
