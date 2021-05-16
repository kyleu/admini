package controller

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vsandbox"

	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app/sandbox"
)

func SandboxList(w http.ResponseWriter, r *http.Request) {
	act("sandbox.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Data = sandbox.AllSandboxes
		return render(r, w, as, &vsandbox.List{}, ps, "sandbox")
	})
}

func SandboxRun(w http.ResponseWriter, r *http.Request) {
	act("sandbox.run", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		sb := sandbox.AllSandboxes.Get(key)
		if sb == nil {
			return ersp("no sandbox with key [" + key + "]")
		}
		ret, err := sb.Run(as, ps.Logger.With(zap.String("sandbox", key)))
		if err != nil {
			return "", err
		}
		ps.Data = ret
		if sb.Key == "testbed" {
			return render(r, w, as, &vsandbox.Testbed{}, ps, "sandbox", sb.Key)
		}
		return render(r, w, as, &vsandbox.Run{Key: key, Title: sb.Title, Result: ret}, ps, "sandbox", sb.Key)
	})
}
