package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/app/workspace"
	"github.com/kyleu/admini/views/vworkspace"
	"net/http"
	"strings"

	"github.com/kyleu/admini/app"
)

func WorkspaceProject(w http.ResponseWriter, r *http.Request) {
	actWorkspace("workspace", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		sourceKey := mux.Vars(r)["key"]

		paths := util.SplitAndTrim(r.URL.Path, "/")
		if len(paths) < 2 {
			return ersp("no source provided in path [%v]", r.URL.Path)
		}
		if paths[0] != "x" {
			return ersp("provided path [%v] is not part of the workspace", r.URL.Path)
		}

		return loadProject(r, w, sourceKey, paths[2:], as, ps)
	})
}

func loadProject(r *http.Request, w http.ResponseWriter, projectKey string, paths []string, as *app.State, ps *cutil.PageState) (string, error) {
	prj, err := currentApp.Projects.Load(projectKey)
	if err != nil {
		return "", fmt.Errorf("unable to load project: %w", err)
	}

	ps.Title = prj.Title
	ps.RootPath = currentApp.Route("workspace", "key", prj.Key)
	ps.RootTitle = prj.Title
	ps.SearchPath = currentApp.Route("search")
	ps.ProfilePath = currentApp.Route("profile")
	ps.Menu = workspace.ProjectMenu(currentApp, prj)

	if len(paths) == 0 {
		return render(r, w, as, &vworkspace.WorkspaceOverview{}, ps)
	}

	return ersp("unhandled [%v] project action [%v]", prj.Key, strings.Join(paths, "/"))
}
