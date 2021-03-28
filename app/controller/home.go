package controller

import (
	"net/http"

	"github.com/kyleu/admini/gen/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	act("home", w, r, func() (string, error) {
		return tmpl(templates.Index(w))
	})
}
