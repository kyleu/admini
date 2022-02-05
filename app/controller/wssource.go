package controller

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/app/workspace"
)

func WorkspaceSource(rc *fasthttp.RequestCtx) {
	actWorkspace("workspace.source", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		sourceKey, err := RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}

		path := string(rc.Request.URI().Path())
		paths := util.StringSplitAndTrim(path, "/")
		if len(paths) < 2 {
			return ersp("no source provided in path [%s]", path)
		}
		if paths[0] != "s" {
			return ersp("provided path [%s] is not part of the source workspace", path)
		}

		pv, err := as.Services.Projects.LoadSourceProject(sourceKey)
		if err != nil {
			return "", errors.Wrapf(err, "error loading source and schema info [%s]", path)
		}

		paths = paths[2:]

		ps.Title = pv.Project.Name()
		ps.RootIcon = pv.Project.IconWithFallback()
		ps.RootPath = fmt.Sprintf("/s/%s", sourceKey)
		ps.RootTitle = pv.Project.Name()
		ps.SearchPath = defaultSearchPath
		ps.ProfilePath = defaultProfilePath

		m, err := workspace.ProjectMenu(as, pv)
		if err != nil {
			return "", errors.Wrapf(err, "error creating menu for project [%s]", pv.Project.Key)
		}
		ps.Menu = m

		a, remaining := pv.Project.Actions.Get(paths)

		wr := &cutil.WorkspaceRequest{
			T: "s", K: sourceKey, RC: rc, PS: ps, Item: a, Path: remaining,
			Project: pv.Project, Sources: pv.Sources, Schemata: pv.Schemata, Context: ps.Context,
		}

		return handleAction(wr, a, rc, as)
	})
}
