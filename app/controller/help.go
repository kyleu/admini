package controller

import (
	"net/http"

	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/views/vhelp"
)

func Feedback(w http.ResponseWriter, r *http.Request) {
	act("feedback", w, r, func(st *ctx.PageState) (string, error) {
		return render(w, &vhelp.Feedback{}, st, "feedback")
	})
}

func Help(w http.ResponseWriter, r *http.Request) {
	act("help", w, r, func(st *ctx.PageState) (string, error) {
		return render(w, &vhelp.Help{}, st, "help")
	})
}
