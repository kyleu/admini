// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"github.com/valyala/fasthttp"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/site"
	"admini.dev/admini/app/util"
)

func Site(rc *fasthttp.RequestCtx) {
	ActSite("site", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		path := util.StringSplitAndTrim(string(rc.Request.URI().Path()), "/")
		redir, page, bc, err := site.Handle(path, rc, as, ps)
		if err != nil {
			return "", err
		}
		if redir != "" {
			return redir, nil
		}
		return Render(rc, as, page, ps, bc...)
	})
}
