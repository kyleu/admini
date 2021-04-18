package controller

import (
	"net/http"

	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/views"
)

func Home(w http.ResponseWriter, r *http.Request) {
	act("home", w, r, func(app *ctx.AppState, page *ctx.PageState) (string, error) {
		return render(w, &views.Home{}, page)
	})
}
