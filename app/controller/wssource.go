package controller

import (
	"github.com/kyleu/admini/app/project"
	"net/http"

	"github.com/kyleu/admini/app/result"
	"github.com/kyleu/admini/views"
	"github.com/kyleu/admini/views/vsource"

	"github.com/pkg/errors"

	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/app/workspace"
	"github.com/kyleu/admini/views/vworkspace"

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
			return ersp("provided path [%v] is not part of the source workspace", r.URL.Path)
		}

		src, sch, err := loadSource(sourceKey)
		if err != nil {
			return "", errors.Wrap(err, "error loading source and schema info ["+r.URL.Path+"]")
		}

		sources := source.Sources{src}

		paths = paths[2:]

		ps.Title = src.Name()
		ps.RootTitle = src.Name()
		ps.RootPath = currentApp.Route(sRoute, "key", src.Key)
		ps.SearchPath = currentApp.Route("search")
		ps.ProfilePath = currentApp.Route("profile")
		ps.Menu = workspace.SourceMenu(currentApp, src.Key, sch)

		if len(paths) == 0 {
			return render(r, w, as, &vworkspace.WorkspaceOverview{}, ps)
		}

		if paths[0] == "_" {
			return sourceAction(r, w, as, ps, src, paths[1:])
		}

		prj := &project.Project{}

		i, remaining := sch.ModelsByPackage().Get(paths)
		wr := &workspaceRequest{T: sRoute, K: src.Key, W: w, R: r, AS: as, PS: ps, Source: src.Key, Item: i, Path: remaining, Sources: sources, Project: prj}
		return handle(wr)
	})
}

func sourceAction(r *http.Request, w http.ResponseWriter, as *app.State, ps *cutil.PageState, src *source.Source, paths []string) (string, error) {
	if len(paths) == 0 {
		return render(r, w, as, &vworkspace.WorkspaceOverview{}, ps)
	}

	switch paths[0] {
	case "sql":
		return actionSQL(r, w, as, ps, src)
	default:
		return render(r, w, as, &views.TODO{Message: "Unhandled source action [" + paths[0] + "]"}, ps, "Not found")
	}
}

func actionSQL(r *http.Request, w http.ResponseWriter, as *app.State, ps *cutil.PageState, src *source.Source) (string, error) {
	sql := r.URL.Query().Get("sql")
	if r.Method == http.MethodPost {
		f, _ := cutil.ParseForm(r)
		x := f.Get("sql")
		if x != nil && x.Value != "" {
			sql = x.Value
		}
	}
	var res *result.Result
	if sql == "" {
		ps.Title = "SQL Playground"
	} else {
		ld, err := as.Loaders.Get(src.Type, src.Key, src.Config)
		if err != nil {
			return "", errors.Wrap(err, "unable to create loader")
		}

		r, err := ld.Query(sql)
		if err != nil {
			return "", errors.Wrap(err, "unable to execute query")
		}
		res = r
		ps.Title = "SQL Result"
		ps.Data = res
	}

	return render(r, w, as, &vsource.SQLPlayground{SQL: sql, Res: res}, ps, "sql")
}

func loadSource(sourceKey string) (*source.Source, *schema.Schema, error) {
	src, err := currentApp.Sources.Load(sourceKey, false)
	if err != nil {
		return nil, nil, err
	}
	sch, err := currentApp.Sources.SchemaFor(src.Key)
	if err != nil {
		return nil, nil, err
	}
	return src, sch, nil
}
