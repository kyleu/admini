package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/app/workspace"
	"github.com/kyleu/admini/views"
	"github.com/kyleu/admini/views/vworkspace"

	"github.com/kyleu/admini/app"
)

func actWorkspace(key string, w http.ResponseWriter, r *http.Request, f func(source string, sch *schema.Schema, as *app.State, ps *cutil.PageState) (string, error)) {
	ps := actPrepare(r, w)
	source := mux.Vars(r)["key"]
	src, err := currentApp.Sources.Load(source)
	if err != nil {
		msg := "error loading source [%v]: %+v"
		http.Error(w, fmt.Sprintf(msg, key, err), http.StatusInternalServerError)
	}
	sch, err := currentApp.Sources.SchemaFor(source)
	if err != nil {
		msg := "error loading schema [%v]: %+v"
		http.Error(w, fmt.Sprintf(msg, key, err), http.StatusInternalServerError)
	}

	ps.RootPath = currentApp.Route("workspace", "key", source)
	ps.RootTitle = src.Title
	ps.SearchPath = currentApp.Route("search")
	ps.ProfilePath = currentApp.Route("profile")
	ps.Menu = workspace.SchemaMenu(currentApp, source, sch)
	actComplete(key, ps, r.URL.Path, w, func() (string, error) { return f(source, sch, currentApp, ps) })
}

func Workspace(w http.ResponseWriter, r *http.Request) {
	actWorkspace("workspace", w, r, func(source string, sch *schema.Schema, as *app.State, ps *cutil.PageState) (string, error) {
		paths := util.SplitAndTrim(r.URL.Path, "/")
		if len(paths) < 2 {
			return ersp("no source provided in path [%v]", r.URL.Path)
		}
		if paths[0] != "x" {
			return ersp("provided path [%v] is not part of the workspace", r.URL.Path)
		}
		paths = paths[2:]
		if len(paths) == 0 {
			return render(r, w, as, &vworkspace.WorkspaceOverview{Schema: sch}, ps)
		}

		i, remaining := sch.ModelsByPackage().Get(paths)

		switch t := i.(type) {
		case *schema.Model:
			return handleModel(w, r, as, ps, t, remaining)
		case *schema.ModelPackage:
			return handlePackage(w, r, as, ps, t, remaining)
		case error:
			return ersp("provided path [%v] can't be loaded: %+v", r.URL.Path, t)
		case nil:
			return ersp("nil path [%v] can't be loaded: %+v", r.URL.Path, t)
		default:
			return ersp("unhandled type: %T", t)
		}
	})
}

func handleModel(w http.ResponseWriter, r *http.Request, as *app.State, ps *cutil.PageState, m *schema.Model, remaining []string) (string, error) {
	if len(remaining) == 0 {
		ps.Data = m
		page := &vworkspace.ModelList{Model: m}
		return render(r, w, as, page, ps, append(m.Pkg, m.Key)...)
	}
	ps.Data = m
	page := &views.TODO{Message: fmt.Sprintf("unhandled model action [%v]", strings.Join(remaining, "/"))}
	return render(r, w, as, page, ps, append(m.Pkg, m.Key)...)
}

func handlePackage(w http.ResponseWriter, r *http.Request, as *app.State, ps *cutil.PageState, mp *schema.ModelPackage, remaining []string) (string, error) {
	if len(remaining) == 0 {
		ps.Data = mp
		page := &vworkspace.PackageDetail{Pkg: mp}
		return render(r, w, as, page, ps, append(mp.Pkg, mp.Key)...)
	}
	ps.Data = mp
	page := &views.TODO{Message: fmt.Sprintf("unhandled package action [%v]", strings.Join(remaining, "/"))}
	return render(r, w, as, page, ps, append(mp.Pkg, mp.Key)...)
}
