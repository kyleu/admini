package controller

import (
	"fmt"
	"net/url"

	theme2 "github.com/kyleu/admini/app/theme"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vuser"
)

func Profile(ctx *fasthttp.RequestCtx) {
	act("profile", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Profile"
		ps.Data = ps.Profile
		theme := theme2.ThemeDefault

		prvs, err := currentApp.Auth.Providers()
		if err != nil {
			return "", errors.Wrap(err, "can't load providers")
		}

		redir := "/"
		ref := string(ctx.Request.Header.Peek("Referer"))
		if ref != "" {
			u, err := url.Parse(ref)
			if err == nil && u != nil {
				redir = u.Path
			}
		}

		return render(ctx, as, &vuser.Profile{Profile: ps.Profile, Theme: theme, Providers: prvs, Referrer: redir}, ps, "Profile")
	})
}

func ProfileSave(ctx *fasthttp.RequestCtx) {
	act("profile.save", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(ctx)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		name := frm["name"]
		mode := frm["mode"]
		theme := frm["theme"]

		println(fmt.Sprintf("PROFILE SAVE: %s, %s, %s", name, mode, theme))

		return flashAndRedir(true, "profile saved", "/profile", ctx, ps)
	})
}
