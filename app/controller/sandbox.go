package controller

import (
	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app/sandbox"
	"net/http"

	"github.com/kyleu/admini/gen/templates"
)

func SandboxList(w http.ResponseWriter, r *http.Request) {
	act("sandbox.list", w, r, func() (string, error) {
		return tmpl(templates.SandboxList(w))
	})
}

func SandboxRun(w http.ResponseWriter, r *http.Request) {
	act("sandbox.run", w, r, func() (string, error) {
		key := mux.Vars(r)["key"]
		sb := sandbox.AllSandboxes.Get(key)
		if sb == nil {
			return ersp("no sandbox with key [" + key + "]")
		}
		return tmpl(templates.SandboxRun(key, w))
	})
}
