package controller

import (
	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/views"
	"github.com/kyleu/admini/views/vprofile"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	act("home", w, r, func(st *ctx.PageState) (string, error) {
		views.WriteRender(w, &vprofile.Profile{Basic: with(st, "Profile")})
		return "", nil
	})
}
