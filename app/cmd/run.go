package cmd

import (
	"fmt"
	"runtime"

	"github.com/fasthttp/router"
	"github.com/pkg/errors"

	"github.com/kirsle/configdir"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/log"
	"github.com/kyleu/admini/app/util"
	"go.uber.org/zap"
)

var AppBuildInfo *app.BuildInfo

func Run() (*zap.SugaredLogger, error) {
	if AppBuildInfo == nil {
		return nil, errors.New("no build info")
	}

	flags := parseFlags()
	if flags.ConfigDir == "" {
		flags.ConfigDir = configdir.LocalConfig(util.AppName)
		_ = configdir.MakePath(flags.ConfigDir)
	}

	logger, err := log.InitLogging(flags.Debug, flags.JSON)
	if err != nil {
		return logger, err
	}

	addr := fmt.Sprintf("%s:%d", flags.Address, flags.Port)
	startLog := logger.With(zap.Bool("debug", flags.Debug), zap.String("address", flags.Address), zap.Int("port", flags.Port))
	startLog.Infof("[%s v%s] %s", util.AppName, AppBuildInfo.Version, util.AppURL)
	startLog.Infof("starting using address [%s] on %s:%s", addr, runtime.GOOS, runtime.GOARCH)

	switch flags.Mode {
	case "app":
		return startApp(flags, logger)
	case "site":
		return startSite(flags, logger)
	default:
		return nil, errors.New("invalid mode [" + flags.Mode + "]")
	}
}

func webserver(addr string, logger *zap.SugaredLogger, r *router.Router) (*zap.SugaredLogger, error) {
	s := &fasthttp.Server{Handler: r.Handler, Name: util.AppName, ReadBufferSize: 32768}
	return logger, s.ListenAndServe(addr)
}
