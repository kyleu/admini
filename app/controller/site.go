package controller

import (
	"github.com/fasthttp/router"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/site"
	"github.com/kyleu/admini/app/util"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app"
)

func SiteRoutes() *router.Router {
	w := fasthttp.CompressHandler
	r := router.New()

	r.GET("/", w(Site))

	r.GET(defaultProfilePath, w(ProfileSite))
	r.POST(defaultProfilePath, w(ProfileSave))
	r.GET("/auth/{key}", w(AuthDetail))
	r.GET("/auth/{key}/callback", w(AuthCallback))
	r.GET("/auth/{key}/logout", w(AuthLogout))

	r.GET("/favicon.ico", Favicon)
	r.GET("/assets/{_:*}", Static)

	r.GET("/{path:*}", w(Site))

	r.OPTIONS("/", w(Options))
	r.OPTIONS("/{_:*}", w(Options))
	r.NotFound = NotFound

	return r
}

func Site(ctx *fasthttp.RequestCtx) {
	actSite("site", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		path := util.SplitAndTrim(string(ctx.Request.URI().Path()), "/")
		redir, page, bc, err := site.Handle(path, ctx, as, ps)
		if err != nil {
			return "", err
		}
		if redir != "" {
			return redir, nil
		}
		return render(ctx, as, page, ps, bc...)
	})
}