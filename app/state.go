package app

import (
	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/source"
	"go.uber.org/zap"
)

type State struct {
	Router       *mux.Router
	Files        filesystem.FileLoader
	Sources      *source.Service
	Projects     *project.Service
	Loaders      *loader.Service
	RootLogger   *zap.SugaredLogger
	routerLogger *zap.SugaredLogger
}

func NewState(r *mux.Router, f *filesystem.FileSystem, ss *source.Service, ps *project.Service, ls *loader.Service, log *zap.SugaredLogger) *State {
	rl := log.With(zap.String("service", "router"))
	return &State{Router: r, Files: f, Sources: ss, Projects: ps, Loaders: ls, RootLogger: log, routerLogger: rl}
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
