package controller

import (
	"net/http"

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
			return ersp("error loading source and schema info [%v]: %w", r.URL.Path, err)
		}

		paths = paths[2:]

		ps.Title = src.Title + " - source"
		ps.RootTitle = src.Title
		ps.RootPath = currentApp.Route(workspaceSourceRoute, "key", src.Key)
		ps.SearchPath = currentApp.Route("search")
		ps.ProfilePath = currentApp.Route("profile")
		ps.Menu = workspace.SourceMenu(currentApp, src.Key, sch)

		if len(paths) == 0 {
			return render(r, w, as, &vworkspace.WorkspaceOverview{}, ps)
		}

		i, remaining := sch.ModelsByPackage().Get(paths)
		wr := &workspaceRequest{T: workspaceSourceRoute, K: src.Key, W: w, R: r, AS: as, PS: ps, I: i, Path: remaining, Src: src}
		return handle(wr)
	})
}

func loadSource(sourceKey string) (*source.Source, *schema.Schema, error) {
	src, err := currentApp.Sources.Load(sourceKey)
	if err != nil {
		return nil, nil, err
	}
	sch, err := currentApp.Sources.SchemaFor(src.Key)
	if err != nil {
		return nil, nil, err
	}
	return src, sch, nil
}
