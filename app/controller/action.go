package controller

import (
	"net/http"

	"github.com/pkg/errors"

	"admini.dev/admini/app"
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/project"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/vproject"
)

func ActionEdit(w http.ResponseWriter, r *http.Request) {
	Act("action.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		p, a, _, err := loadAction(r, as, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "error loading project and action")
		}
		ps.Title = a.Name()
		ps.Data = a
		page := &vproject.ActionEdit{Project: p, Act: a}
		return Render(w, r, as, page, ps, append([]string{"projects", p.Key}, a.Path()...)...)
	})
}

func loadAction(r *http.Request, as *app.State, logger util.Logger) (*project.Project, *action.Action, []string, error) {
	key, err := cutil.PathString(r, "key", false)
	if err != nil {
		return nil, nil, nil, err
	}
	p, err := as.Services.Projects.Load(key, false, logger)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "unable to load project [%s]", key)
	}

	path, err := cutil.PathString(r, "path", false)
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
