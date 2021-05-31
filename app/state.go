package app

import (
	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/project/action"
	"github.com/kyleu/admini/app/source"
	"go.uber.org/zap"
)

type State struct {
	Debug        bool
	Router       *mux.Router
	Files        filesystem.FileLoader
	Sources      *source.Service
	Projects     *project.Service
	Loaders      *loader.Service
	RootLogger   *zap.SugaredLogger
	routerLogger *zap.SugaredLogger
}

func NewState(debug bool, r *mux.Router, f *filesystem.FileSystem, ls *loader.Service, log *zap.SugaredLogger) (*State, error) {
	rl := log.With(zap.String("service", "router"))
	ss := source.NewService(action.TypeSource.Key, f, ls, log)
	ps := project.NewService("project", f, ss, ls, log)

	ret := &State{Debug: debug, Router: r, Files: f, Sources: ss, Projects: ps, Loaders: ls, RootLogger: log, routerLogger: rl}
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
