package app

import (
	"fmt"

	"github.com/fasthttp/router"
	"github.com/kyleu/admini/app/auth"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/app/util"
	"go.uber.org/zap"
)

type BuildInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

func (b *BuildInfo) String() string {
	return fmt.Sprintf("%s v%s: commit %s on %s", util.AppName, b.Version, b.Commit, b.Date)
}

type State struct {
	Debug        bool
	BuildInfo    *BuildInfo
	Router       *router.Router
	Files        filesystem.FileLoader
	Auth         *auth.Service
	Sources      *source.Service
	Projects     *project.Service
	Loaders      *loader.Service
	RootLogger   *zap.SugaredLogger
	routerLogger *zap.SugaredLogger
}

func NewState(debug bool, bi *BuildInfo, r *router.Router, f filesystem.FileLoader, ls *loader.Service, log *zap.SugaredLogger) (*State, error) {
	rl := log.With(zap.String("service", "router"))
	as := auth.NewService("", log)
	ss := source.NewService(f, ls, log)
	ps := project.NewService(f, ss, ls, log)

	ret := &State{
		Debug:        debug,
		BuildInfo:    bi,
		Router:       r,
		Files:        f,
		Auth:         as,
		Sources:      ss,
		Projects:     ps,
		Loaders:      ls,
		RootLogger:   log,
		routerLogger: rl,
	}
	return ret, nil
}
