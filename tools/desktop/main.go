package main

import (
	"fmt"

	"github.com/kyleu/admini/app/cmd"
	"github.com/kyleu/admini/app/util"
	"github.com/webview/webview"
)

func main() {
	if err := run(true); err != nil {
		panic(err)
	}
}

func run(debug bool) error {
	port, err := cmd.Lib()
	if err != nil {
		return err
	}
	loadWebview(debug, int(port))
	return nil
}

func loadWebview(debug bool, port int) {
	w := webview.New(debug)
	defer w.Destroy()

	w.SetTitle(util.AppName)
	w.SetSize(1280, 720, webview.HintNone)
	w.Navigate(fmt.Sprintf("http://localhost:%d", port))

	w.Run()
}
