package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
)

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
		if a.Key != newKey {
			na, _ := p.Actions.Get(a.Pkg.With(newKey))
			if na != nil {
				return "", errors.Errorf("Action with key [%s] already exists in package [%s]", newKey, a.Pkg.String())
			}
			a.Key = newKey
		}

		if a.TypeKey == action.TypeSeparator.Key {
			a.Title = frm.GetStringOpt("title")
			a.Description = frm.GetStringOpt("description")
			icon := frm.GetStringOpt("icon")
			if icon != "" {
				a.Icon = icon
			}
		}
		switch a.TypeKey {
		case action.TypeStatic.Key:
			format := frm.GetStringOpt("format")
			if format == "" {
				format = "text"
			}
			a.Config["format"] = format
			a.Config["content"] = frm.GetStringOpt("content")
		}

		// p.Actions.Set(a)

		err = as.Services.Projects.Save(p, true)
		if err != nil {
			return "", err
		}

		return flashAndRedir(true, "saved action", fmt.Sprintf("/project/%s", p.Key), rc, ps)
	})
}
