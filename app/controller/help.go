package controller

import (
	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/views"
	"github.com/kyleu/admini/views/vhelp"
	"net/http"
)

func Feedback(w http.ResponseWriter, r *http.Request) {
	act("feedback", w, r, func(st *ctx.PageState) (string, error) {
		views.WriteRender(w, &vhelp.Feedback{Basic: with(st, "feedback")})
		return "", nil
	})
}

func Help(w http.ResponseWriter, r *http.Request) {
	act("help", w, r, func(st *ctx.PageState) (string, error) {
		views.WriteRender(w, &vhelp.Help{Basic: with(st, "help")})
		return "", nil
	})
}
