package controller

import (
	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/views/vsandbox"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app/sandbox"
)

func SandboxList(w http.ResponseWriter, r *http.Request) {
	act("sandbox.list", w, r, func(st *ctx.PageState) (string, error) {
		return render(w, &vsandbox.SandboxList{}, st, "sandbox")
	})
}

func SandboxRun(w http.ResponseWriter, r *http.Request) {
	act("sandbox.run", w, r, func(st *ctx.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		sb := sandbox.AllSandboxes.Get(key)
		if sb == nil {
			return ersp("no sandbox with key [" + key + "]")
		}
		ret, err := sb.Run()
		if err != nil {
			return "", err
		}
		return render(w, &vsandbox.SandboxRun{Key: key, Title: sb.Title, Result: ret}, st, "sandbox", sb.Key)
	})
}
