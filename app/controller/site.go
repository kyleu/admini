package controller

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views"
)

func SiteRoutes() *router.Router {
	w := fasthttp.CompressHandler
	r := router.New()
	r.GET("/", w(Site))
	r.GET("/{path:*}", w(Site))

	r.GET("/favicon.ico", Favicon)
	r.GET("/assets/{_:*}", Static)

	r.OPTIONS("/", w(Options))
	r.OPTIONS("/{_:*}", w(Options))
	r.NotFound = NotFound

	return r
}

func Site(ctx *fasthttp.RequestCtx) {
	act("site", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		return render(ctx, as, &views.Debug{}, ps)
	})
}
