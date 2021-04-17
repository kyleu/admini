package controller

import (
	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/views"
	"github.com/kyleu/admini/views/vsearch"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	act("home", w, r, func(st *ctx.PageState) (string, error) {
		q := r.URL.Query().Get("q")
		views.WriteRender(w, &vsearch.SearchResults{Basic: with(st, "search"), Q: q})
		return "", nil
	})
}
