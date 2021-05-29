package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kyleu/admini/views"
	"github.com/pkg/errors"

	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/app/workspace"
)

const sRoute = "workspace.source"

func WorkspaceSource(w http.ResponseWriter, r *http.Request) {
	actWorkspace(sRoute, w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		sourceKey := mux.Vars(r)["key"]

		paths := util.SplitAndTrim(r.URL.Path, "/")
		if len(paths) < 2 {
			return ersp("no source provided in path [%v]", r.URL.Path)
		}
		if paths[0] != "s" {
			return ersp("provided path [%v] is not part of the source workspace", r.URL.Path)
		}

		pv, err := currentApp.Projects.LoadSourceProject(sourceKey)
		if err != nil {
			return "", errors.Wrap(err, "error loading source and schema info ["+r.URL.Path+"]")
		}

		paths = paths[2:]

		ps.Title = pv.Project.Name()
		ps.RootPath = currentApp.Route(sRoute, "key", sourceKey)
		ps.RootTitle = pv.Project.Name()
		ps.SearchPath = currentApp.Route("search")
		ps.ProfilePath = currentApp.Route("profile")

		m, err := workspace.ProjectMenu(currentApp, pv)
		if err != nil {
			return "", errors.Wrap(err, "error creating menu for project ["+pv.Project.Key+"]")
		}
		ps.Menu = m

		a, remaining := pv.Project.Actions.Get(paths)

		wr := &cutil.WorkspaceRequest{T: sRoute, K: sourceKey, W: w, R: r, AS: as, PS: ps, Item: a, Path: remaining, Project: pv.Project, Sources: pv.Sources, Schemata: pv.Schemata}

		return handleAction(wr, a)
	})
}

func whoops(req *cutil.WorkspaceRequest, msg string, path ...string) (string, error) {
	page := &views.TODO{Message: fmt.Sprintf("%v [%v]", msg, strings.Join(req.Path, "/"))}
	return renderWS(req, page, path...)
}
