// Package controller $PF_IGNORE$
package controller

import (
	"net/url"

	"github.com/valyala/fasthttp"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views"
)

var homeContent = util.ValueMap{
	"_": util.AppName,
	"urls": map[string]string{
		"projects":  "/project",
		"sources":   "/source",
		"sandboxes": "/sandbox",
	},
}

func Home(rc *fasthttp.RequestCtx) {
	Act("home", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		projects, _ := as.Services.Projects.List(ps.Context, ps.Logger)
		sources, _ := as.Services.Sources.List(ps.Logger)
		ps.Data = homeContent
		return Render(rc, as, &views.Home{Sources: sources, Projects: projects}, ps)
	})
}

func Refresh(rc *fasthttp.RequestCtx) {
	Act("refresh", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		redir := "/"
		ref := string(rc.Request.Header.Peek("Referer"))
		if ref != "" {
			u, err := url.Parse(ref)
			if err == nil && u != nil {
				redir = u.Path
			}
		}
		as.Themes.Clear()
		as.Services.Loaders.Clear()
		as.Services.Sources.Clear()
		as.Services.Projects.Clear()
		const msg = "Cleared all caches"
		return FlashAndRedir(true, msg, redir, rc, ps)
	})
}
