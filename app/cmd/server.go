package cmd

import (
	"fmt"
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
	return http.ListenAndServe(fmt.Sprintf("%s:%v", address, port), r)
}
