// Code generated by Project Forge, see https://projectforge.dev for details.
package controller

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/views"
)

func About(ctx *fasthttp.RequestCtx) {
	act("about", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Data = util.AppName + " v" + as.BuildInfo.Version
		return render(ctx, as, &views.About{}, ps)
	})
}
