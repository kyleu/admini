package cmd

import (
	"fmt"
	"net"
	"strconv"

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
	return Start(parseFlags())
}

func Start(flags *Flags) (*zap.SugaredLogger, error) {
	l, err := rootLogger(flags)
	if err != nil {
		return nil, err
	}

	startLog := l.With(zap.Bool("debug", flags.Debug), zap.String("address", flags.Address), zap.Int("port", int(flags.Port)))
	startLog.Infof("[%s v%s] %s", util.AppName, AppBuildInfo.Version, util.AppURL)

	switch flags.Mode {
	case "server":
		return startServer(flags, l)
	case "site":
		return startSite(flags, l)
	case "all":
		go func() {
			f := flags.Clone(flags.Port + 1)
			_, err := startSite(f, l)
			if err != nil {
				l.Errorf("can't start marketing site: %+v", err)
			}
		}()
		return startServer(flags, l)
	default:
		return nil, errors.New("invalid mode [" + flags.Mode + "]")
	}
}

func rootLogger(flags *Flags) (*zap.SugaredLogger, error) {
	if AppBuildInfo == nil {
		return nil, errors.New("no build info")
	}

	if flags.ConfigDir == "" {
		flags.ConfigDir = configdir.LocalConfig(util.AppName)
		_ = configdir.MakePath(flags.ConfigDir)
	}

	l, err := log.InitLogging(flags.Debug, flags.JSON)
	if err != nil {
		return l, err
	}
	return l, nil
}

func listen(address string, port uint16) (uint16, net.Listener, error) {
	l, err := net.Listen("tcp", fmt.Sprintf("%v:%v", address, port))
	if err != nil {
		return port, nil, errors.Wrap(err, fmt.Sprintf("unable to listen on port [%v]", port))
	}
	if port == 0 {
		addr := l.Addr().String()
		_, portStr := util.SplitStringLast(addr, ':', true)
		actualPort, err := strconv.Atoi(portStr)
		if err != nil {
			return 0, nil, errors.Wrap(err, "invalid port ["+portStr+"]")
		}
		port = uint16(actualPort)
	}
	return port, l, nil
}

func serve(name string, listener net.Listener, r *router.Router) error {
	err := fasthttp.Serve(listener, r.Handler)
	if err != nil {
		return errors.Wrap(err, "unable to run http server")
	}
	return nil
}

func listenandserve(name string, addr string, port uint16, r *router.Router) (uint16, error) {
	p, l, err := listen(addr, port)
	if err != nil {
		return p, err
	}
	err = serve(name, l, r)
	if err != nil {
		return p, err
	}
	return 0, nil
}
