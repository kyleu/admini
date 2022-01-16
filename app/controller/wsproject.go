package controller

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/lib/telemetry"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/app/workspace"
)

func actWorkspace(key string, rc *fasthttp.RequestCtx, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	as := _currentAppState
	ps := loadPageState(rc, key, as)
	actComplete(key, as, ps, rc, f)
}

func WorkspaceProject(rc *fasthttp.RequestCtx) {
	actWorkspace("workspace", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		projectKey, err := RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}

		paths := util.StringSplitAndTrim(string(rc.Path()), "/")
		if len(paths) < 2 {
			return ersp("no source provided in path [%s]", string(rc.Path()))
		}
		if paths[0] != "x" {
			return ersp("provided path [%s] is not part of the project workspace", string(rc.Path()))
		}

		pv, err := as.Services.Projects.LoadView(projectKey)
		if err != nil {
			return "", errors.Wrapf(err, "error loading project [%s]", projectKey)
		}

		paths = paths[2:]

		ps.Title = pv.Project.Name()
		ps.RootIcon = pv.Project.IconWithFallback()
		ps.RootPath = fmt.Sprintf("/x/%s", pv.Project.Key)
		ps.RootTitle = pv.Project.Name()
		ps.SearchPath = defaultSearchPath
		ps.ProfilePath = defaultProfilePath

		m, err := workspace.ProjectMenu(as, pv)
		if err != nil {
			return "", errors.Wrapf(err, "error creating menu for project [%s]", pv.Project.Key)
		}
		ps.Menu = m

		a, remaining := pv.Project.Actions.Get(paths)

		ps.Context, _ = telemetry.StartSpan(ps.Context, "workspace", strings.Join(paths, "/"))
		wr := &cutil.WorkspaceRequest{
			T: "x", K: pv.Project.Key, Ctx: rc, PS: ps, Item: a, Path: remaining,
			Project: pv.Project, Sources: pv.Sources, Schemata: pv.Schemata, Context: ps.Context,
		}
		return handleAction(wr, a, rc, as)
	})
}

func handleAction(req *cutil.WorkspaceRequest, act *action.Action, rc *fasthttp.RequestCtx, as *app.State) (string, error) {
	if req == nil {
		return "", errors.New("nil project request")
	}
	if act == nil {
		act = &action.Action{}
	}
	res, err := workspace.ActionHandler(req, act, as)
	if err != nil {
		return "", err
	}

	if res.Redirect != "" {
		return flashAndRedir(true, res.Title, res.Redirect, rc, req.PS)
	}

	req.PS.Title = res.Title
	req.PS.Data = res.Data

	return render(req.Ctx, as, res.Page, req.PS, res.Breadcrumbs...)
}
