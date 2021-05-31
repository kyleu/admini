package controller

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kyleu/admini/views/verror"

	"github.com/pkg/errors"

	"github.com/gorilla/sessions"
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/views"
	"github.com/kyleu/admini/views/layout"
)

var (
	currentApp   *app.State
	initialIcons = []string{"search"}
)

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

func render(r *http.Request, w http.ResponseWriter, appState *app.State, page layout.Page, pageState *cutil.PageState, breadcrumbs ...string) (string, error) {
	println(len(breadcrumbs), strings.Join(breadcrumbs, "/"))

	pageState.Breadcrumbs = append(pageState.Breadcrumbs, breadcrumbs...)
	ct := cutil.GetContentType(r)
	if pageState.Data != nil {
		if cutil.IsContentTypeJSON(ct) {
			return cutil.RespondJSON(w, "", pageState.Data)
		} else if cutil.IsContentTypeXML(ct) {
			return cutil.RespondXML(w, "", pageState.Data)
		}
	}
	startNanos := time.Now().UnixNano()
	views.WriteRender(w, page, appState, pageState)
	pageState.RenderElapsed = float64((time.Now().UnixNano()-startNanos)/int64(time.Microsecond)) / float64(1000)
	return "", nil
}

func renderWS(req *cutil.WorkspaceRequest, page layout.Page, bc ...string) (string, error) {
	return render(req.R, req.W, req.AS, page, req.PS, bc...)
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

func flashError(err error, redir string, w http.ResponseWriter, r *http.Request, ps *cutil.PageState) (string, error) {
	return flashAndRedir(false, err.Error(), redir, w, r, ps)
}

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
