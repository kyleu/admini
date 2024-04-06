package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"admini.dev/admini/app"
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/telemetry"
	"admini.dev/admini/app/util"
	"admini.dev/admini/app/workspace"
)

func actWorkspace(key string, w http.ResponseWriter, r *http.Request, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	as := _currentAppState
	ps := cutil.LoadPageState(as, w, r, key, _currentAppRootLogger)
	actComplete(key, as, ps, w, r, f)
}

func WorkspaceProject(w http.ResponseWriter, r *http.Request) {
	actWorkspace("workspace", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		projectKey, err := cutil.PathString(r, "key", false)
		if err != nil {
			return "", err
		}

		paths := util.StringSplitAndTrim(r.URL.Path, "/")
		if len(paths) < 2 {
			return ERsp("no project provided in path [%s]", r.URL.Path)
		}
		if paths[0] != "x" {
			return ERsp("provided path [%s] is not part of the project workspace", r.URL.Path)
		}

		pv, err := as.Services.Projects.LoadView(projectKey, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load project [%s]", projectKey)
		}

		paths = paths[2:]

		ps.Title = pv.Project.Name()
		ps.RootIcon = pv.Project.IconWithFallback()
		ps.RootPath = fmt.Sprintf("/x/%s", pv.Project.Key)
		ps.RootTitle = pv.Project.Name()
		ps.SearchPath = cutil.DefaultSearchPath
		ps.ProfilePath = cutil.DefaultProfilePath

		m, err := workspace.ProjectMenu(as, pv)
		if err != nil {
			return "", errors.Wrapf(err, "error creating menu for project [%s]", pv.Project.Key)
		}
		ps.Menu = m

		a, remaining := pv.Project.Actions.Get(paths)

		ctx, span, logger := telemetry.StartSpan(ps.Context, "workspace:"+strings.Join(paths, "/"), ps.Logger)
		defer span.Complete()
		ps.Context = ctx
		ps.Logger = logger
		wr := &cutil.WorkspaceRequest{
			T: "x", K: pv.Project.Key, Req: r, ReqBody: ps.RequestBody, Rsp: w, PS: ps, Item: a, Path: remaining,
			Project: pv.Project, Sources: pv.Sources, Schemata: pv.Schemata, Context: ps.Context,
		}
		return handleAction(wr, a, as)
	})
}

func handleAction(req *cutil.WorkspaceRequest, act *action.Action, as *app.State) (string, error) {
	if req == nil {
		return "", errors.New("nil project request")
	}
	if act == nil {
		act = action.RootAction
	}
	res, err := workspace.ActionHandler(req, act, as)
	if err != nil {
		return "", err
	}

	if res.Redirect != "" {
		return FlashAndRedir(true, res.Title, res.Redirect, req.Rsp, req.PS)
	}

	req.PS.Title = res.Title
	req.PS.Data = res.Data
	req.PS.SearchPath = req.Route("search")

	return Render(req.Rsp, req.Req, as, res.Page, req.PS, res.Breadcrumbs...)
}
