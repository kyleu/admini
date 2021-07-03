package controller

import (
	"fmt"

	"github.com/kyleu/admini/app/qualify"

	"github.com/kyleu/admini/app/action"

	"github.com/kyleu/admini/app/util"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vproject"
)

func ProjectList(ctx *fasthttp.RequestCtx) {
	act("project.list", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		p, err := as.Projects.List()
		if err != nil {
			return "", errors.Wrap(err, "unable to load project list")
		}
		ps.Title = "Projects"
		ps.Data = p
		return render(ctx, as, &vproject.List{Projects: p}, ps, "projects")
	})
}

func ProjectDetail(ctx *fasthttp.RequestCtx) {
	act("project.detail", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		prj, err := as.Projects.LoadView(key)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load project [%s]", key)
		}
		ps.Title = prj.Project.Name()
		ps.Data = prj.Project
		return render(ctx, as, &vproject.Detail{View: prj}, ps, "projects", prj.Project.Key)
	})
}

func ProjectTest(ctx *fasthttp.RequestCtx) {
	act("project.test", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		v, err := as.Projects.LoadView(key)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load project [%s]", key)
		}

		req := qualify.NewRequest(action.TypeModel.Key, "list", action.TypeSource.Key, "admini_test", action.TypeModel.Key, "public/simple")

		q, err := qualify.Qualify(req, v.Project.Actions, v.Schemata)
		if err != nil {
			return "", errors.Wrapf(err, "unable to qualify project [%s]", key)
		}

		ps.Title = v.Project.Name() + " - Test"
		ps.Data = util.ValueMap{
			"request": req,
			"qualify": q,
		}

		view := &vproject.Test{Message: fmt.Sprintf("Project [%s]: OK", v.Project.Key)}
		return render(ctx, as, view, ps, "projects", v.Project.Key, "test")
	})
}
