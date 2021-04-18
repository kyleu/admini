package controller

import (
	"net/http"

	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/views/vsettings"
)

func Settings(w http.ResponseWriter, r *http.Request) {
	act("settings", w, r, func(app *ctx.AppState, page *ctx.PageState) (string, error) {
		return render(w, &vsettings.Settings{}, page, "settings")
	})
}
