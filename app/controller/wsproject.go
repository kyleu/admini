package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/app/workspace"
	"github.com/kyleu/admini/views/vworkspace"

	"github.com/kyleu/admini/app"
)

func WorkspaceProject(w http.ResponseWriter, r *http.Request) {
	actWorkspace("workspace", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		projectKey := mux.Vars(r)["key"]

		paths := util.SplitAndTrim(r.URL.Path, "/")
		if len(paths) < 2 {
			return ersp("no source provided in path [%v]", r.URL.Path)
		}
		if paths[0] != "x" {
			return ersp("provided path [%v] is not part of the project workspace", r.URL.Path)
		}

		prj, err := currentApp.Projects.Load(projectKey)
		if err != nil {
			return ersp("error loading project [%v]: %w", projectKey, err)
		}

		paths = paths[2:]

		ps.Title = prj.Title
		ps.RootPath = currentApp.Route("workspace", "key", prj.Key)
		ps.RootTitle = prj.Title
		ps.SearchPath = currentApp.Route("search")
		ps.ProfilePath = currentApp.Route("profile")
		ps.Menu = workspace.ProjectMenu(currentApp, prj)

		if len(paths) == 0 {
			return render(r, w, as, &vworkspace.WorkspaceOverview{}, ps)
		}

		action, remaining := prj.Actions.Get(paths)

		wr := &workspaceRequest{T: workspaceProjectRoute, K: prj.Key, W: w, R: r, AS: as, PS: ps, I: action, Path: remaining, Src: nil}
		return handle(wr)
	})
}
