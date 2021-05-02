package app

import (
	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/app/util"
)

type State struct {
	Router  *mux.Router
	Files   filesystem.FileLoader
	Sources *source.Service
	Projects *project.Service
	Loaders *loader.Service
}

func (a *State) Route(act string, pairs ...string) string {
	route := a.Router.Get(act)
	if route == nil {
		util.LogWarn("cannot find route at path [" + act + "]")
		return "/route/notfound/" + act
	}
	u, err := route.URL(pairs...)
	if err != nil {
		util.LogWarn("cannot bind route at path [" + act + "]")
		return "/route/error/" + act
	}
	return u.Path
}
