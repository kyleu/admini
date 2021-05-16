package cmd

import (
	"fmt"
	"net/http"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/loader/lmock"
	"github.com/kyleu/admini/app/loader/lpostgres"
	"github.com/kyleu/admini/app/log"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
	"go.uber.org/zap"
)

func Run() error {
	logger, err := log.InitLogging(true)
	if err != nil {
		return err
	}

	address := ""

	logger.With(zap.String("address", address), zap.Int("port", int(util.AppPort))).Info("[" + util.AppName + "]")

	r, err := controller.BuildRouter()
	if err != nil {
		return err
	}

	f := filesystem.NewFileSystem("data", logger)
	ls := loader.NewService()
	ls.Set(schema.OriginPostgres, lpostgres.NewLoader(logger))
	ls.Set(schema.OriginMock, lmock.NewLoader(logger))

	st, err := app.NewState(r, f, ls, logger)
	if err != nil {
		return err
	}
	controller.SetState(st)

	return http.ListenAndServe(fmt.Sprintf("%s:%v", address, util.AppPort), r)
}
