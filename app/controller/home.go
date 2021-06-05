package controller

import (
	"github.com/kyleu/admini/app/util"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views"
)

var homeContent = util.ValueMap{
	"_": util.AppName,
	"urls": map[string]string{
		"projects":  "/project",
		"sources":   "/source",
		"sandboxes": "/sandbox",
	},
}

func Home(ctx *fasthttp.RequestCtx) {
	act("home", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		projects, _ := as.Projects.List()
		sources, _ := as.Sources.List()
		ps.Data = homeContent
		return render(ctx, as, &views.Home{Sources: sources, Projects: projects}, ps)
	})
}
