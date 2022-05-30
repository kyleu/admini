package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"admini.dev/admini/app"
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/qualify"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/vproject"
)

func ProjectList(rc *fasthttp.RequestCtx) {
	act("project.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		p, err := as.Services.Projects.List(ps.Context, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to load project list")
		}
		ps.Title = "Projects"
		ps.Data = p
		return render(rc, as, &vproject.List{Projects: p}, ps, "projects")
	})
}

func ProjectDetail(rc *fasthttp.RequestCtx) {
	act("project.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		prj, err := as.Services.Projects.LoadView(key, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load project [%s]", key)
		}
		ps.Title = prj.Project.Name()
		ps.Data = prj.Project
		return render(rc, as, &vproject.Detail{View: prj}, ps, "projects", prj.Project.Key)
	})
}

func ProjectTest(rc *fasthttp.RequestCtx) {
	act("project.test", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		v, err := as.Services.Projects.LoadView(key, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load project [%s]", key)
		}

		req := qualify.NewRequest(action.TypeModel.Key, "list", action.TypeSource.Key, "admini_test", action.TypeModel.Key, "public/simple")

		q, err := qualify.Qualify(req, v.Project.Actions, v.Schemata)
		if err != nil {
			return "", errors.Wrapf(err, "unable to qualify project [%s]", key)
		}

		var models []any
		for _, s := range v.Schemata {
			for _, m := range s.Models {
				if true {
					models = append(models, m.Key)
				}
			}
		}

		ps.Title = v.Project.Name() + " - Test"
		ps.Data = util.ValueMap{
			"request": req,
			"qualify": q,
			"models":  models,
		}

		view := &vproject.Test{Message: fmt.Sprintf("Project [%s]: OK", v.Project.Key)}
		return render(rc, as, view, ps, "projects", v.Project.Key, "test")
	})
}
