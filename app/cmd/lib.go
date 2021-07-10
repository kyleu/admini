package cmd

import (
	"fmt"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
)

// Lib starts the application as a library, returning the actual TCP port the server is listening on (as an int32 to make interop easier)
func Lib() (int, error) {
	if AppBuildInfo == nil {
		AppBuildInfo = &app.BuildInfo{Version: "TODO", Commit: "TODO", Date: "TODO"}
	}
	f := &Flags{Address: "0.0.0.0", Port: 0, Mode: "app"}

	logger, err := rootLogger(f)
	if err != nil {
		return 0, err
	}

	r, logger, err := loadApp(f, logger)
	if err != nil {
		return 0, err
	}

	port, listener, err := listen(f.Address, f.Port)
	if err != nil {
		return 0, err
	}

	logger.Info(fmt.Sprintf("%v library started on port [%v]", util.AppName, port))

	go func() {
		e := serve("app", listener, r)
		if e != nil {
			panic(errors.WithStack(e))
		}
	}()

	return int(port), nil
}
