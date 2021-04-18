package controller

import (
	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/views/vsettings"
	"net/http"
)

func Settings(w http.ResponseWriter, r *http.Request) {
	act("settings", w, r, func(st *ctx.PageState) (string, error) {
		return render(w, &vsettings.Settings{}, st, "settings")
	})
}
