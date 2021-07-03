package controller

import (
	"net/url"

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

func Refresh(ctx *fasthttp.RequestCtx) {
	act("refresh", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		redir := "/"
		ref := string(ctx.Request.Header.Peek("Referer"))
		if ref != "" {
			u, err := url.Parse(ref)
			if err == nil && u != nil {
				redir = u.Path
			}
		}
		currentApp.Loaders.Clear()
		currentApp.Sources.Clear()
		currentApp.Projects.Clear()
		msg := "Cleared all caches"
		return flashAndRedir(true, msg, redir, ctx, ps)
	})
}
