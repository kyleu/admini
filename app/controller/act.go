package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kyleu/admini/app/menu"

	"go.uber.org/zap"

	"github.com/kyleu/admini/views/verror"

	"github.com/gorilla/sessions"
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/util"
)

func act(key string, w http.ResponseWriter, r *http.Request, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	ps := actPrepare(r, w)
	clean(ps)
	actComplete(key, ps, r, w, f)
}

func actWorkspace(key string, w http.ResponseWriter, r *http.Request, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	ps := actPrepare(r, w)
	actComplete(key, ps, r, w, f)
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

func actComplete(key string, ps *cutil.PageState, r *http.Request, w http.ResponseWriter, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	status := http.StatusOK
	cutil.WriteCORS(w)
	startNanos := time.Now().UnixNano()
	redir, err := f(currentApp, ps)
	if err != nil {
		status = http.StatusInternalServerError
		w.WriteHeader(status)

		ps.Logger.Errorf("error running action [%v]: %+v", key, err)

		if len(ps.Breadcrumbs) == 0 {
			ps.Breadcrumbs = []string{"Error"}
		}
		errDetail := util.GetErrorDetail(err)
		page := &verror.Error{Err: errDetail}

		clean(ps)
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
	l.Debugf("processed request in [%.3fms] (render: %.3fms)", elapsedMillis, ps.RenderElapsed)
}

func clean(ps *cutil.PageState) {
	if ps.RootPath == "" {
		ps.RootPath = currentApp.Route("home")
	}
	if ps.RootTitle == "" {
		ps.RootTitle = util.AppName
	}
	if ps.SearchPath == "" {
		ps.SearchPath = currentApp.Route("search")
	}
	if ps.ProfilePath == "" {
		ps.ProfilePath = currentApp.Route("profile")
	}
	if len(ps.Menu) == 0 {
		ps.Menu = menu.For(currentApp)
	}
}
