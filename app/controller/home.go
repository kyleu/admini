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

func Settings(w http.ResponseWriter, r *http.Request) {
	act("settings", w, r, func() (string, error) {
		return tmpl(templates.Settings(w))
	})
}

func Feedback(w http.ResponseWriter, r *http.Request) {
	act("feedback", w, r, func() (string, error) {
		return tmpl(templates.Feedback(w))
	})
}

func Help(w http.ResponseWriter, r *http.Request) {
	act("help", w, r, func() (string, error) {
		return tmpl(templates.Help(w))
	})
}
