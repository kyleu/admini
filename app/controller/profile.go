package controller

import (
	"github.com/kyleu/admini/app/user"
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
		theme := user.ThemeDefault

		prvs, err := currentApp.Auth.Providers()
		if err != nil {
			return "", errors.Wrap(err, "can't load providers")
		}

		return render(ctx, as, &vuser.Profile{Profile: ps.Profile, Theme: theme, Providers: prvs}, ps, "Profile")
	})
}
