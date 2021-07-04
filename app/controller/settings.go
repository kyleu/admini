package controller

import (
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/theme"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/views"
	"github.com/kyleu/admini/views/vtheme"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vsettings"
)

func Settings(ctx *fasthttp.RequestCtx) {
	act("settings", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		settings := &project.Settings{Test: "OK"}
		ps.Title = "Settings"
		ps.Data = settings
		return render(ctx, as, &vsettings.Settings{Settings: settings}, ps, "settings")
	})
}

func ThemeEdit(ctx *fasthttp.RequestCtx) {
	act("theme.edit", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		var t *theme.Theme
		if key == "new" {
			t = theme.ThemeDefault.Clone(key)
		} else {
			t = as.Themes.Get(key)
		}
		if t == nil {
			return "", errors.Wrap(err, "no theme found with key ["+key+"]")
		}
		ps.Data = t
		ps.Title = "Edit theme"
		page := &vtheme.Edit{Theme: t}
		return render(ctx, as, page, ps, "settings", "Themes", t.Key)
	})
}

func ThemeSave(ctx *fasthttp.RequestCtx) {
	act("theme.save", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		frm, err := cutil.ParseForm(ctx)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		orig := as.Themes.Get(key)

		newKey, err := frm.GetString("key", false)
		if err != nil {
			return "", err
		}
		if newKey == "new" {
			newKey = util.RandomString(12)
		}

		l := orig.Light.Clone().ApplyMap(frm, "light-")
		d := orig.Dark.Clone().ApplyMap(frm, "dark-")

		t := &theme.Theme{Key: newKey, Light: l, Dark: d}

		err = as.Themes.Save(t)
		if err != nil {
			return "", errors.Wrap(err, "unable to save theme")
		}

		ps.Data = map[string]interface{}{"key": newKey, "new": t, "orig": orig, "form": frm}
		ps.Title = "Edit theme"
		page := &views.Debug{}
		return render(ctx, as, page, ps)
	})
}
