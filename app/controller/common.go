package controller

import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/views/verror"
	"github.com/kyleu/admini/views/vhelp"
	"github.com/pkg/errors"
	"net/http"
	"runtime/debug"
)

func Options(w http.ResponseWriter, r *http.Request) {
	cutil.WriteCORS(w)
	w.WriteHeader(http.StatusOK)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	act("notfound", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		cutil.WriteCORS(w)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		ps.Logger.Warnf("%v %v returned [%d]", r.Method, r.URL.Path, http.StatusNotFound)
		if ps.Title == "" {
			ps.Title = "404"
		}
		ps.Data = "404 not found"
		return render(r, w, as, &verror.NotFound{}, ps, "Not Found")
	})
}

func Modules(w http.ResponseWriter, r *http.Request) {
	act("modules", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		mods, ok := debug.ReadBuildInfo()
		if !ok {
			return "", errors.New("unable to gather modules")
		}
		ps.Title = "Modules"
		ps.Data = mods.Deps
		return render(r, w, as, &vhelp.Modules{Mods: mods.Deps}, ps, "modules")
	})
}

func Routes(w http.ResponseWriter, r *http.Request) {
	act("routes", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		routes := cutil.ExtractRoutes(as.Router)
		ps.Title = "Routes"
		ps.Data = routes
		return render(r, w, as, &vhelp.Routes{Routes: routes}, ps, "routes")
	})
}
