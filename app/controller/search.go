package controller

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/search"
	"github.com/kyleu/admini/views/vsearch"
)

func Search(rc *fasthttp.RequestCtx) {
	act("search", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := string(rc.URI().QueryArgs().Peek("q"))
		params := &search.Params{Q: q}
		results, errs := search.Search(ps.Context, as, params)
		ps.Title = "Search Results"
		ps.Data = results
		return render(rc, as, &vsearch.Results{Params: params, Results: results, Errors: errs}, ps, "Search")
	})
}
