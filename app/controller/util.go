package controller

import (
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/kyleu/admini/app/util"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/views"
	"github.com/kyleu/admini/views/layout"
	"github.com/kyleu/admini/views/vhelp"
)

var currentApp *app.State

var initialIcons = []string{"app", "search", "profile"}

var sessionKey = func() string {
	x := os.Getenv("SESSION_KEY")
	if x == "" {
		x = "random_secret_key"
	}
	return x
}()

var store = sessions.NewCookieStore([]byte(sessionKey))

func SetState(a *app.State) {
	currentApp = a
}

func render(r *http.Request, w http.ResponseWriter, appState *app.State, page layout.Page, pageState *cutil.PageState, bc ...string) (string, error) {
	pageState.Breadcrumbs = bc
	ct := getContentType(r)
	if pageState.Data != nil && isContentTypeJSON(ct) {
		return respondJSON(w, "", pageState.Data)
	}
	views.WriteRender(w, page, appState, pageState)
	return "", nil
}

func ersp(msg string, args ...interface{}) (string, error) {
	return "", fmt.Errorf(msg, args...)
}

func flashAndRedir(success bool, msg string, redir string, w http.ResponseWriter, r *http.Request, ps *cutil.PageState) (string, error) {
	status := "error"
	if success {
		status = "success"
	}
	ps.Session.AddFlash(status + ":" + msg)
	_ = ps.Session.Save(r, w)
	if strings.HasPrefix(redir, "/") {
		return redir, nil
	}
	if strings.HasPrefix(redir, "http") {
		util.LogWarn("flash redirect attempted for non-local request")
		return "/", nil
	}
	return redir, nil
}

func Options(w http.ResponseWriter, r *http.Request) {
	writeCORS(w)
	w.WriteHeader(http.StatusOK)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	act("notfound", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		writeCORS(w)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		util.LogInfo("%v %v returned [%d]", r.Method, r.URL.Path, http.StatusNotFound)
		ps.Data = "404 not found"
		return render(r, w, as, &views.NotFound{}, ps, "Not Found")
	})
}

func Modules(w http.ResponseWriter, r *http.Request) {
	act("modules", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		mods, ok := debug.ReadBuildInfo()
		if !ok {
			return "", fmt.Errorf("unable to gather modules")
		}
		ps.Data = mods.Deps
		return render(r, w, as, &vhelp.Modules{Mods: mods.Deps}, ps, "modules")
	})
}

func Routes(w http.ResponseWriter, r *http.Request) {
	act("routes", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		routes := cutil.ExtractRoutes(as.Router)
		ps.Data = routes
		return render(r, w, as, &vhelp.Routes{Routes: routes}, ps, "routes")
	})
}
