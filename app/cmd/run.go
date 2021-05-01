package cmd

import (
	"strings"

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
