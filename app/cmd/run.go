package cmd

import (
	"fmt"

	"github.com/kyleu/admini/app/loader/lsqlite"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/loader/lmock"
	"github.com/kyleu/admini/app/loader/lpostgres"
	"github.com/kyleu/admini/app/log"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

type Flags struct {
	Address string
	Debug   bool
	Profile bool
}

func Run() error {
	logger, err := log.InitLogging(true)
	if err != nil {
		return err
	}

	flags := parseFlags()

	logger.With(
		zap.Bool("debug", flags.Debug),
		zap.String("address", flags.Address),
		zap.Int("port", util.AppPort),
	).Infof("[%s]", util.AppName)

	r := controller.BuildRouter()

	f := filesystem.NewFileSystem("data", logger)
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
	pflag.StringVar(&ret.Address, "addr", "127.0.0.1", "address to listen on, defaults to [127.0.0.1]")
	pflag.BoolVar(&ret.Debug, "debug", false, "enables verbose logging and additional checks")
	pflag.Parse()
	return ret
}
