package main

import (
	"os"
	"strings"

	"github.com/kyleu/admini/app/cmd"
	"github.com/kyleu/admini/app/util"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"server"}
	}
	run(args[0], args[1:])
}

func run(mode string, args []string) {
	util.LogInfo("[%v] starting in [%v] mode with args [%v]", util.AppName, mode, strings.Join(args, ", "))
	var err error
	switch mode {
	case "server":
		err = cmd.StartServer("", util.AppPort)
	default:
		util.LogInfo(util.AppName + "!")
	}
	if err != nil {
		util.LogInfo("%+v", err)
	}
}
