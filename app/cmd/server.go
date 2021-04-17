package cmd

import (
	"fmt"
	"github.com/kyleu/admini/app/ctx"
	"net/http"

	"github.com/kyleu/admini/app/controller"
	"github.com/kyleu/admini/app/util"
)

func StartServer(address string, port uint16) error {
	util.LogInfo("server starting on [%s:%v]", address, port)
	r, err := controller.BuildRouter()
	if err != nil {
		return err
	}
	ctx.ActiveRouter = r
	return http.ListenAndServe(fmt.Sprintf("%s:%v", address, port), r)
}
