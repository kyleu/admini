package controller

import (
	"github.com/kyleu/admini/app/project"
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
