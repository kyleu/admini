// Package controller $PF_IGNORE$
package controller

import (
	"net/url"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/util"
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

func Home(rc *fasthttp.RequestCtx) {
	act("home", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		projects, _ := as.Services.Projects.List(ps.Context)
		sources, _ := as.Services.Sources.List()
		ps.Data = homeContent
		return render(rc, as, &views.Home{Sources: sources, Projects: projects}, ps)
	})
}

func Refresh(rc *fasthttp.RequestCtx) {
	act("refresh", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
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
		msg := "Cleared all caches"
		return flashAndRedir(true, msg, redir, rc, ps)
	})
}
