// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"github.com/valyala/fasthttp"

	"admini.dev/app"
	"admini.dev/app/controller/cutil"
	"admini.dev/app/util"
	"admini.dev/views"
)

func About(rc *fasthttp.RequestCtx) {
	act("about", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Data = util.AppName + " v" + as.BuildInfo.Version
		return render(rc, as, &views.About{}, ps, "about")
	})
}
