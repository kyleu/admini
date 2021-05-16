package app

import (
	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/app/watcher"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type State struct {
	Router       *mux.Router
	Files        filesystem.FileLoader
	Sources      *source.Service
	Projects     *project.Service
	Loaders      *loader.Service
	Watcher      *watcher.Service
	RootLogger   *zap.SugaredLogger
	routerLogger *zap.SugaredLogger
}

func NewState(r *mux.Router, f *filesystem.FileSystem, ls *loader.Service, log *zap.SugaredLogger) (*State, error) {
	rl := log.With(zap.String("service", "router"))
	ss := source.NewService("source", f, ls, log)
	ps := project.NewService("project", f, ss, ls, log)
	ws, err := watcher.NewService(f.Root(), log)
	if err != nil {
		return nil, errors.Wrap(err, "error initializing file watcher")
	}

	ret := &State{Router: r, Files: f, Sources: ss, Projects: ps, Loaders: ls, Watcher: ws, RootLogger: log, routerLogger: rl}
	return ret, nil
}

func (a *State) Route(act string, pairs ...string) string {
	route := a.Router.Get(act)
	if route == nil {
		a.routerLogger.Warnf("cannot find route at path [%v]", act)
		return "/route/notfound/" + act
	}
	u, err := route.URL(pairs...)
	if err != nil {
		a.routerLogger.Warnf("cannot bind route at path [%v]", act)
		return "/route/error/" + act
	}
	return u.Path
}
