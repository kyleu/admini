package controller

import (
	"net/http"

	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/views/vhelp"
)

func Feedback(w http.ResponseWriter, r *http.Request) {
	act("feedback", w, r, func(app *ctx.AppState, page *ctx.PageState) (string, error) {
		return render(w, &vhelp.Feedback{}, page, "feedback")
	})
}

func Help(w http.ResponseWriter, r *http.Request) {
	act("help", w, r, func(app *ctx.AppState, page *ctx.PageState) (string, error) {
		return render(w, &vhelp.Help{}, page, "help")
	})
}
