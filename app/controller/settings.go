package controller

import (
	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/views"
	"github.com/kyleu/admini/views/vsettings"
	"net/http"
)

func Settings(w http.ResponseWriter, r *http.Request) {
	act("settings", w, r, func(st *ctx.PageState) (string, error) {
		views.WriteRender(w, &vsettings.Settings{Basic: with(st, "settings")})
		return "", nil
	})
}
