package controller

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/kyleu/admini/views/verror"

	"github.com/gorilla/sessions"
	"github.com/kyleu/admini/app/menu"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/util"
)

func act(key string, w http.ResponseWriter, r *http.Request, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	ps := actPrepare(r, w)
	ps.RootPath = currentApp.Route("home")
	ps.RootTitle = util.AppName
	ps.SearchPath = currentApp.Route("search")
	ps.ProfilePath = currentApp.Route("profile")
	ps.Menu = menu.For(currentApp)
	actComplete(key, ps, r, w, func() (string, error) { return f(currentApp, ps) })
}

func actWorkspace(key string, w http.ResponseWriter, r *http.Request, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	ps := actPrepare(r, w)
	actComplete(key, ps, r, w, func() (string, error) { return f(currentApp, ps) })
}

func actPrepare(r *http.Request, w http.ResponseWriter) *cutil.PageState {
	logger := currentApp.RootLogger.With(zap.String("path", r.URL.Path))

	session, err := store.Get(r, util.AppKey)
	if err != nil {
		logger.Warnf("error retrieving session: %+v", err)
	}
	if session.IsNew {
		session.Options = &sessions.Options{Path: "/", HttpOnly: true, SameSite: http.SameSiteDefaultMode}
		err = session.Save(r, w)
		if err != nil {
			logger.Warnf("can't save session: %+v", err)
		}
	}

	flashes := util.StringArrayFromInterfaces(session.Flashes())
	if len(flashes) > 0 {
		err = session.Save(r, w)
		if err != nil {
			logger.Warnf("cannot save session flashes: %+v", err)
		}
	}

	return &cutil.PageState{Method: r.Method, URL: r.URL, Flashes: flashes, Session: session, Icons: initialIcons, Logger: logger}
}

func actComplete(key string, ps *cutil.PageState, r *http.Request, w http.ResponseWriter, f func() (string, error)) {
	startNanos := time.Now().UnixNano()
	writeCORS(w)
	redir, err := f()
	status := http.StatusOK
	if err != nil {
		status = http.StatusInternalServerError
		w.WriteHeader(status)

		ps.Logger.Errorf("error running action [%v]: %+v", key, err)

		if len(ps.Breadcrumbs) == 0 {
			ps.Breadcrumbs = []string{"Error"}
		}
		errDetail := util.GetErrorDetail(err)
		page := &verror.Error{Err: errDetail}
		redir, err = render(r, w, currentApp, page, ps)
		if err != nil {
			_, _ = w.Write([]byte(fmt.Sprintf("error while running error handler: %+v", err)))
		}
	}
	if redir != "" {
		w.Header().Set("Location", redir)
		status = http.StatusFound
		w.WriteHeader(status)
	}
	elapsedMillis := float64((time.Now().UnixNano()-startNanos)/int64(time.Microsecond)) / float64(1000)
	l := ps.Logger.With(zap.String("method", ps.Method), zap.Int("status", status), zap.Float64("elapsed", elapsedMillis))
	l.Debugf("processed request in [%.3fms]", elapsedMillis)
}
