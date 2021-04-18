package controller

import (
	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/views/vsource"
	"net/http"

	"github.com/gorilla/mux"
)

func SourceList(w http.ResponseWriter, r *http.Request) {
	act("source.list", w, r, func(st *ctx.PageState) (string, error) {
		return render(w, &vsource.SourceList{}, st, "sources")
	})
}

func SourceDetail(w http.ResponseWriter, r *http.Request) {
	act("source.detail", w, r, func(st *ctx.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		d := &source.Source{Key: key}
		return render(w, &vsource.SourceDetail{Source: d}, st, "sources", d.Key)
	})
}
