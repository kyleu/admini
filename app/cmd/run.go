package cmd

import (
	"fmt"
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
	"net/http"
)

type Flags struct {
	Address string
	Debug   bool
	Profile bool
}

func (f *Flags) String() string {
	return fmt.Sprintf("addr: %v, debug: %v, profile: %v", f.Address, f.Debug, f.Profile)
}

func Run() error {
	logger, err := log.InitLogging(true)
	if err != nil {
		return err
	}

	flags := parseFlags()

	logger.With(zap.String("address", flags.Address), zap.Int("port", int(util.AppPort))).Info("[" + util.AppName + "]")

	r, err := controller.BuildRouter()
	if err != nil {
		return err
	}

	f := filesystem.NewFileSystem("data", logger)
	ls := loader.NewService()
	ls.Set(schema.OriginPostgres, lpostgres.NewLoader(logger, flags.Debug))
	ls.Set(schema.OriginMock, lmock.NewLoader(logger))

	st, err := app.NewState(flags.Debug, r, f, ls, logger)
	if err != nil {
		return err
	}
	controller.SetState(st)

	return http.ListenAndServe(fmt.Sprintf("%s:%v", flags.Address, util.AppPort), r)
}

func parseFlags() *Flags {
	ret := &Flags{}
	pflag.StringVar(&ret.Address, "addr", "127.0.0.1", "address to listen on, defaults to [127.0.0.1]")
	pflag.BoolVar(&ret.Debug, "debug", false, "enables verbose logging and additional checks")
	pflag.Parse()
	return ret
}
