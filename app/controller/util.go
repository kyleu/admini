package controller

import (
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/gorilla/sessions"
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
	pageState.Breadcrumbs = append(pageState.Breadcrumbs, bc...)
	ct := getContentType(r)
	if pageState.Data != nil && isContentTypeJSON(ct) {
		return respondJSON(w, "", pageState.Data)
	}
	startNanos := time.Now().UnixNano()
	views.WriteRender(w, page, appState, pageState)
	pageState.RenderElapsed = float64((time.Now().UnixNano()-startNanos)/int64(time.Microsecond)) / float64(1000)
	return "", nil
}

func ersp(msg string, args ...interface{}) (string, error) {
	return "", errors.Errorf(msg, args...)
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
		ps.Logger.Warn("flash redirect attempted for non-local request")
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
		ps.Logger.Warnf("%v %v returned [%d]", r.Method, r.URL.Path, http.StatusNotFound)
		ps.Data = "404 not found"
		return render(r, w, as, &views.NotFound{}, ps, "Not Found")
	})
}

func Modules(w http.ResponseWriter, r *http.Request) {
	act("modules", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		mods, ok := debug.ReadBuildInfo()
		if !ok {
			return "", errors.New("unable to gather modules")
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
