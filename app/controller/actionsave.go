package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"admini.dev/admini/app"
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
)

func ActionSave(rc *fasthttp.RequestCtx) {
	Act("action.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		p, a, _, err := loadAction(rc, as, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "error loading project and action")
		}

		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}

		newKey := frm.GetStringOpt("key")
		if a.Key != newKey {
			na, _ := p.Actions.Get(a.Pkg.With(newKey))
			if na != nil {
				return "", errors.Errorf("Action with key [%s] already exists in package [%s]", newKey, a.Pkg.String())
			}
			a.Key = newKey
		}

		title := frm.GetStringOpt("title")
		if title == "" {
			title = a.Title
		}
		if title == "" {
			title = "-"
		}
		a.Title = title
		a.Description = frm.GetStringOpt("description")
		a.Icon = frm.GetStringOpt("icon")

		if a.TypeKey == action.TypeStatic.Key {
			format := frm.GetStringOpt("format")
			if format == "" {
				format = "text"
			}
			a.Config["format"] = format
			a.Config["content"] = frm.GetStringOpt("content")
		}

		acts := p.Actions.Clone()
		// p.Actions.Set(a)
		p.Actions = acts

		err = as.Services.Projects.Save(p, true, ps.Logger)
		if err != nil {
			return "", err
		}

		return FlashAndRedir(true, "saved action", fmt.Sprintf("/project/%s", p.Key), rc, ps)
	})
}
