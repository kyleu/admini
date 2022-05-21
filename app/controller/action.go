package controller

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"admini.dev/admini/app"
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/project"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/vproject"
)

func ActionEdit(rc *fasthttp.RequestCtx) {
	act("action.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		p, a, _, err := loadAction(rc, as)
		if err != nil {
			return "", errors.Wrap(err, "error loading project and action")
		}
		ps.Title = a.Name()
		ps.Data = a
		page := &vproject.ActionEdit{Project: p, Act: a}
		return render(rc, as, page, ps, append([]string{"projects", p.Key}, a.Path()...)...)
	})
}

func loadAction(rc *fasthttp.RequestCtx, as *app.State) (*project.Project, *action.Action, []string, error) {
	key, err := cutil.RCRequiredString(rc, "key", false)
	if err != nil {
		return nil, nil, nil, err
	}
	p, err := as.Services.Projects.Load(key, false)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "unable to load project [%s]", key)
	}

	path, err := cutil.RCRequiredString(rc, "path", false)
	if err != nil {
		return nil, nil, nil, err
	}
	pkg := util.StringSplitAndTrim(path, "/")

	a, remaining := p.Actions.Get(pkg)
	if a == nil {
		return nil, nil, nil, errors.Errorf("no action available at [%s]", path)
	}
	if a.Config == nil {
		a.Config = util.ValueMap{}
	}

	return p, a, remaining, nil
}
