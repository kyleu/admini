package controller

import (
	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/views/vsearch"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	act("home", w, r, func(st *ctx.PageState) (string, error) {
		q := r.URL.Query().Get("q")
		return render(w, &vsearch.SearchResults{Q: q}, st, "search")
	})
}
