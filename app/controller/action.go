package controller

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/kyleu/admini/app/action"

	"github.com/kyleu/admini/views/vproject"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
)

func ActionOrdering(rc *fasthttp.RequestCtx) {
	act("action.ordering", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := rcRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		prj, err := as.Services.Projects.LoadRequired(key, false)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load project [%s]", key)
		}

		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}
		ordering := frm.GetStringOpt("ordering")

		actOrders := action.Orderings{}
		err = util.FromJSON([]byte(ordering), &actOrders)
		if err != nil {
			return "", err
		}
		startNanos := time.Now().UnixNano()
		newActs, err := action.ReorderActions(prj.Actions, actOrders)
		if err != nil {
			return "", err
		}

		count, err := action.SaveAll(prj.Key, newActs, as.Files)
		if err != nil {
			return "", err
		}
		elapsedMillis := float64((time.Now().UnixNano()-startNanos)/int64(time.Microsecond)) / float64(1000)

		_, err = as.Services.Projects.LoadRequired(prj.Key, true)
		if err != nil {
			return "", err
		}
		msg := fmt.Sprintf("saved [%d] %s in [%.3fms]", count, util.Plural(count, "action"), elapsedMillis)
		return flashAndRedir(true, msg, fmt.Sprintf("/project/%s", key), rc, ps)
	})
}

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

func ActionSave(rc *fasthttp.RequestCtx) {
	act("action.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		p, a, _, err := loadAction(rc, as)
		if err != nil {
			return "", errors.Wrap(err, "error loading project and action")
		}

		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}

		newKey := frm.GetStringOpt("key")
		shouldReload := false
		if a.Key != newKey {
			na, _ := p.Actions.Get(a.Pkg.With(newKey))
			if na != nil {
				return "", errors.Errorf("Action with key [%s] already exists in package [%s]", newKey, a.Pkg.String())
			}
			err = as.Services.Projects.DeleteAction(p.Key, a)
			if err != nil {
				return "", err
			}
			a.Key = newKey
			shouldReload = true
		}

		if a.Type != action.TypeSeparator {
			a.Title = frm.GetStringOpt("title")
			a.Description = frm.GetStringOpt("description")
			icon := frm.GetStringOpt("icon")
			if icon != "" {
				a.Icon = icon
			}
		}
		actPath := filepath.Join("project", p.Key, "actions", strings.Join(a.Pkg, "/"))
		_, err = action.Save(actPath, a, as.Files)
		if err != nil {
			return "", err
		}

		if shouldReload {
			err = as.Services.Projects.ReloadProject(p.Key)
			if err != nil {
				return "", err
			}
		}

		return flashAndRedir(true, "saved action", fmt.Sprintf("/project/%s", p.Key), rc, ps)
	})
}

func loadAction(rc *fasthttp.RequestCtx, as *app.State) (*project.Project, *action.Action, []string, error) {
	key, err := rcRequiredString(rc, "key", false)
	if err != nil {
		return nil, nil, nil, err
	}
	p, err := as.Services.Projects.Load(key, false)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "unable to load project [%s]", key)
	}

	path, err := rcRequiredString(rc, "path", false)
	if err != nil {
		return nil, nil, nil, err
	}
	pkg := util.SplitAndTrim(path, "/")

	a, remaining := p.Actions.Get(pkg)
	if a == nil {
		return nil, nil, nil, errors.Errorf("no action available at [%s]", path)
	}

	return p, a, remaining, nil
}
