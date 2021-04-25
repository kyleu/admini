package controller

import (
	"fmt"
	"net/http"
	"time"

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
	actComplete(key, ps, r.URL.Path, w, func() (string, error) { return f(currentApp, ps) })
}

func actPrepare(r *http.Request, w http.ResponseWriter) *cutil.PageState {
	session, err := store.Get(r, util.AppKey)
	if err != nil {
		util.LogWarn(fmt.Sprintf("error retrieving session: %+v", err))
	}
	if session.IsNew {
		session.Options = &sessions.Options{Path: "/", HttpOnly: true, SameSite: http.SameSiteDefaultMode}
		err = session.Save(r, w)
		if err != nil {
			util.LogWarn(fmt.Sprintf("cannot save session: %+v", err))
		}
	}

	flashes := make([]string, 0)
	for _, f := range session.Flashes() {
		flashes = append(flashes, fmt.Sprint(f))
	}

	if len(flashes) > 0 {
		err = session.Save(r, w)
		if err != nil {
			util.LogWarn(fmt.Sprintf("cannot save session: %+v", err))
		}
	}

	return &cutil.PageState{Flashes: flashes, Session: session, Icons: initialIcons}
}

func actComplete(key string, p *cutil.PageState, path string, w http.ResponseWriter, f func() (string, error)) {
	startNanos := time.Now().UnixNano()
	writeCORS(w)
	redir, err := f()
	if err != nil {
		msg := "error running action [%v]: %+v"
		util.LogWarn(msg, key, err)
		http.Error(w, fmt.Sprintf(msg, key, err), http.StatusInternalServerError)
	}
	if redir != "" {
		w.Header().Set("Location", redir)
		w.WriteHeader(http.StatusFound)
	}
	elapsedMillis := float64((time.Now().UnixNano()-startNanos)/int64(time.Microsecond)) / float64(1000)
	util.LogInfo("processed [%v] in [%.3fms]", path, elapsedMillis)
}
