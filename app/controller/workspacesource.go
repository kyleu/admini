package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/app/workspace"
	"github.com/kyleu/admini/views/vworkspace"
	"net/http"

	"github.com/kyleu/admini/app"
)

func WorkspaceSource(w http.ResponseWriter, r *http.Request) {
	actWorkspace("workspace.source", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		sourceKey := mux.Vars(r)["key"]

		paths := util.SplitAndTrim(r.URL.Path, "/")
		if len(paths) < 2 {
			return ersp("no source provided in path [%v]", r.URL.Path)
		}
		if paths[0] != "s" {
			return ersp("provided path [%v] is not part of the workspace", r.URL.Path)
		}

		return loadSource(r, w, sourceKey, paths[2:], as, ps)
	})
}

func loadSource(r *http.Request, w http.ResponseWriter, sourceKey string, paths []string, as *app.State, ps *cutil.PageState) (string, error) {
	src, err := currentApp.Sources.Load(sourceKey)
	if err != nil {
		msg := "error loading source [%v]: %+v"
		http.Error(w, fmt.Sprintf(msg, sourceKey, err), http.StatusInternalServerError)
	}
	sch, err := currentApp.Sources.SchemaFor(src.Key)
	if err != nil {
		msg := "error loading schema [%v]: %+v"
		http.Error(w, fmt.Sprintf(msg, src.Key, err), http.StatusInternalServerError)
	}

	ps.Title = src.Title + " - source"
	ps.RootPath = currentApp.Route("workspace.source", "key", src.Key)
	ps.RootTitle = src.Title
	ps.SearchPath = currentApp.Route("search")
	ps.ProfilePath = currentApp.Route("profile")
	ps.Menu = workspace.SourceMenu(currentApp, src.Key, sch)

	if len(paths) == 0 {
		return render(r, w, as, &vworkspace.WorkspaceOverview{}, ps)
	}

	i, remaining := sch.ModelsByPackage().Get(paths)

	switch t := i.(type) {
	case *schema.Model:
		return handleModel(w, r, as, ps, src, t, remaining)
	case *schema.ModelPackage:
		return handlePackage(w, r, as, ps, src, t, remaining)
	case error:
		return ersp("provided path [%v] can't be loaded: %+v", r.URL.Path, t)
	case nil:
		return ersp("nil path [%v] can't be loaded: %+v", r.URL.Path, t)
	default:
		return ersp("unhandled type: %T", t)
	}
}
