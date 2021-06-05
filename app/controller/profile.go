package controller

import (
	"github.com/kyleu/admini/app/user"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vprofile"
)

func Profile(ctx *fasthttp.RequestCtx) {
	act("profile", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		profile := &user.Profile{Name: "Test"}
		ps.Title = "Profile"
		ps.Data = profile
		return render(ctx, as, &vprofile.Profile{Profile: profile, Theme: user.ThemeDefault}, ps, "Profile")
	})
}
