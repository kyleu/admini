package app

import (
	"fmt"

	"github.com/fasthttp/router"
	"github.com/kyleu/admini/app/auth"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/app/theme"
	"github.com/kyleu/admini/app/util"
	"go.uber.org/zap"
)

type BuildInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

func (b *BuildInfo) String() string {
	if b.Date == "unknown" {
	} else {
		d, _ := util.TimeFromJS(b.Date)
		return fmt.Sprintf("%s (%s)", b.Version, util.TimeToYMD(d))
	}
	return b.Version
}

type State struct {
	Debug        bool
	BuildInfo    *BuildInfo
	Router       *router.Router
	Files        filesystem.FileLoader
	Auth         *auth.Service
	Themes       *theme.Service
	Sources      *source.Service
	Projects     *project.Service
	Loaders      *loader.Service
	routerLogger *zap.SugaredLogger
}

func NewState(debug bool, bi *BuildInfo, r *router.Router, f filesystem.FileLoader, ls *loader.Service, log *zap.SugaredLogger) (*State, error) {
	rl := log.With(zap.String("service", "router"))
	ss := source.NewService(f, ls, log)
	ret := &State{
		Debug:        debug,
		BuildInfo:    bi,
		Router:       r,
		Files:        f,
		Auth:         auth.NewService("", log),
		Themes:       theme.NewService(f, log),
		Sources:      ss,
		Projects:     project.NewService(f, ss, ls, log),
		Loaders:      ls,
		routerLogger: rl,
	}
	return ret, nil
}
