package controller

import (
	"net/http"

	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/views/vsandbox"

	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app/sandbox"
)

func SandboxList(w http.ResponseWriter, r *http.Request) {
	act("sandbox.list", w, r, func(app *ctx.AppState, page *ctx.PageState) (string, error) {
		return render(w, &vsandbox.SandboxList{}, page, "sandbox")
	})
}

func SandboxRun(w http.ResponseWriter, r *http.Request) {
	act("sandbox.run", w, r, func(app *ctx.AppState, page *ctx.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		sb := sandbox.AllSandboxes.Get(key)
		if sb == nil {
			return ersp("no sandbox with key [" + key + "]")
		}
		ret, err := sb.Run()
		if err != nil {
			return "", err
		}
		return render(w, &vsandbox.SandboxRun{Key: key, Title: sb.Title, Result: ret}, page, "sandbox", sb.Key)
	})
}
