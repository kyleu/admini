package controller

import (
	"net/http"
	"strings"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vsearch"
)

func Search(w http.ResponseWriter, r *http.Request) {
	act("search", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := r.URL.Query().Get("q")
		q = strings.TrimSpace(q)
		results := []string{"a", "b", "c"}
		ps.Data = results
		return render(r, w, as, &vsearch.Results{Q: q, Results: results}, ps, "Search")
	})
}
