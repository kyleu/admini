package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kyleu/admini/app/controller"
	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/app/util"
)

func Run(mode string, args []string) {
	util.LogInfo("[%v] starting in [%v] mode with args [%v]", util.AppName, mode, strings.Join(args, ", "))
	var err error
	switch mode {
	case "server":
		err = StartServer("", util.AppPort)
	default:
		util.LogInfo(util.AppName + "!")
	}
	if err != nil {
		util.LogInfo("%+v", err)
	}
}

func StartServer(address string, port uint16) error {
	util.LogInfo("server starting on [%s:%v]", address, port)

	r, err := controller.BuildRouter()
	if err != nil {
		return err
	}

	f := filesystem.NewFileSystem("data")
	ds := source.NewService("source", f)

	ctx.App = &ctx.AppState{Router: r, Files: f, Sources: ds}
	return http.ListenAndServe(fmt.Sprintf("%s:%v", address, port), r)
}
