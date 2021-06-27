package cmd

import (
	"fmt"

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
}

func Run() error {
	logger, err := log.InitLogging(true)
	if err != nil {
		return err
	}

	flags := parseFlags()
	if flags.ConfigDir == "" {
		flags.ConfigDir = configdir.LocalConfig(util.AppName)
		_ = configdir.MakePath(flags.ConfigDir)
	}

	logger.With(
		zap.Bool("debug", flags.Debug),
		zap.String("address", flags.Address),
		zap.Int("port", flags.Port),
	).Infof("[%s]", util.AppName)

	r := controller.BuildRouter()

	f := filesystem.NewFileSystem(flags.ConfigDir, logger)
	ls := loader.NewService()
	ls.Set(schema.OriginPostgres, lpostgres.NewLoader(logger))
	ls.Set(schema.OriginSQLite, lsqlite.NewLoader(logger))
	ls.Set(schema.OriginMock, lmock.NewLoader(logger))

	st, err := app.NewState(flags.Debug, r, f, ls, logger)
	if err != nil {
		return err
	}
	controller.SetState(st)

	s := &fasthttp.Server{Handler: r.Handler, Name: util.AppName, ReadBufferSize: 32768}
	return s.ListenAndServe(fmt.Sprintf("%s:%d", flags.Address, util.AppPort))
}

func parseFlags() *Flags {
	ret := &Flags{}
	pflag.StringVarP(&ret.Address, "addr", "a", "127.0.0.1", "address to listen on, defaults to [127.0.0.1]")
	pflag.IntVarP(&ret.Port, "port", "p", util.AppPort, fmt.Sprintf("port to listen on, defaults to [%d]", util.AppPort))
	pflag.StringVarP(&ret.ConfigDir, "dir", "d", "", "directory for configuration, defaults to [~/???]")
	pflag.BoolVarP(&ret.Debug, "verbose", "v", false, "enables verbose logging and additional checks")
	pflag.Parse()
	return ret
}
