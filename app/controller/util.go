package controller

import (
	"errors"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/views"
	"github.com/kyleu/admini/views/layout"
	"github.com/kyleu/admini/views/vhelp"
	"log"
	"net/http"
	"runtime/debug"
)

func with(st *ctx.PageState, bc ...string) layout.Basic {
	st.Breadcrumbs = bc
	return layout.Basic{State: st}
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
		views.WriteRender(w, &views.NotFound{Basic: with(st)})
		return "", nil
	})
}

func Modules(w http.ResponseWriter, r *http.Request) {
	act("home", w, r, func(st *ctx.PageState) (string, error) {
		mods, ok := debug.ReadBuildInfo()
		if !ok {
			return "", errors.New("unable to gather modules")
		}
		views.WriteRender(w, &vhelp.Modules{Basic: with(st, "modules"), Mods: mods.Deps})
		return "", nil
	})
}

func Routes(w http.ResponseWriter, r *http.Request) {
	act("home", w, r, func(st *ctx.PageState) (string, error) {
		views.WriteRender(w, &vhelp.Routes{Basic: with(st, "routes"), Routes: cutil.ExtractRoutes(st.Router)})
		return "", nil
	})
}
