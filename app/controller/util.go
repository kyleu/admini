package controller

import (
	"errors"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/views"
	"github.com/kyleu/admini/views/layout"
	"github.com/kyleu/admini/views/vhelp"
)

func render(w http.ResponseWriter, p layout.Page, st *ctx.PageState, bc ...string) (string, error) {
	st.Breadcrumbs = bc
	views.WriteRender(w, p, st)
	return "", nil
}

func Options(w http.ResponseWriter, r *http.Request) {
	writeCORS(w)
	w.WriteHeader(http.StatusOK)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	act("home", w, r, func(st *ctx.PageState) (string, error) {
		writeCORS(w)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		log.Printf("%v %v returned [%d]", r.Method, r.URL.Path, http.StatusNotFound)
		return render(w, &views.NotFound{}, st, "Not Found")
	})
}

func Modules(w http.ResponseWriter, r *http.Request) {
	act("home", w, r, func(st *ctx.PageState) (string, error) {
		mods, ok := debug.ReadBuildInfo()
		if !ok {
			return "", errors.New("unable to gather modules")
		}
		return render(w, &vhelp.Modules{Mods: mods.Deps}, st, "modules")
	})
}

func Routes(w http.ResponseWriter, r *http.Request) {
	act("home", w, r, func(st *ctx.PageState) (string, error) {
		return render(w, &vhelp.Routes{Routes: cutil.ExtractRoutes(ctx.App.Router)}, st, "routes")
	})
}
