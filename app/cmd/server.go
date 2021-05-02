package cmd

import (
	"fmt"
	"net/http"

	"github.com/kyleu/admini/app/project"

	"github.com/kyleu/admini/app/loader/lmock"
	"github.com/kyleu/admini/app/loader/lpostgres"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/schema"

	"github.com/kyleu/admini/app/controller"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/app/util"
)

func StartServer(address string, port uint16) error {
	util.LogInfo("server starting on [%s:%v]", address, port)

	r, err := controller.BuildRouter()
	if err != nil {
		return err
	}

	f := filesystem.NewFileSystem("data")
	ls := loader.NewService()
	ls.Set(schema.OriginPostgres, lpostgres.NewLoader())
	ls.Set(schema.OriginMock, lmock.NewLoader())
	ss := source.NewService("source", f, ls)
	ps := project.NewService("project", f, ss, ls)

	State := &app.State{Router: r, Files: f, Sources: ss, Projects: ps, Loaders: ls}
	controller.SetState(State)

	return http.ListenAndServe(fmt.Sprintf("%s:%v", address, port), r)
}
