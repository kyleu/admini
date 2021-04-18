package controller

import (
	"net/http"

	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/views/vsource"

	"github.com/gorilla/mux"
)

func SourceList(w http.ResponseWriter, r *http.Request) {
	act("source.list", w, r, func(app *ctx.AppState, page *ctx.PageState) (string, error) {
		return render(w, &vsource.SourceList{}, page, "sources")
	})
}

func SourceDetail(w http.ResponseWriter, r *http.Request) {
	act("source.detail", w, r, func(app *ctx.AppState, page *ctx.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		s, err := app.Sources.Load(key)
		if err != nil {
			return ersp("unable to load source [" + key + "]")
		}
		return render(w, &vsource.SourceDetail{Source: s}, page, "sources", s.Key)
	})
}
