package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/kyleu/admini/app/ctx"
	"github.com/kyleu/admini/app/menu"

	"github.com/kyleu/admini/app/util"
)

var currentApp *ctx.AppState

func SetAppState(a *ctx.AppState) {
	currentApp = a
}

func act(key string, w http.ResponseWriter, r *http.Request, f func(app *ctx.AppState, page *ctx.PageState) (string, error)) {
	page := &ctx.PageState{Menu: menu.For(currentApp.Sources)}
	startNanos := time.Now().UnixNano()
	writeCORS(w)
	redir, err := f(currentApp, page)
	if err != nil {
		msg := "error running action [%v]: %+v"
		util.LogWarn(msg, key, err)
		http.Error(w, fmt.Sprintf(msg, key, err), http.StatusInternalServerError)
	}
	if redir != "" {
		w.Header().Set("Location", redir)
		w.WriteHeader(http.StatusFound)
	}
	elapsedMicros := float64((time.Now().UnixNano()-startNanos)/int64(time.Microsecond)) / float64(1000)
	util.LogInfo("processed [%v] in [%.3fms]", r.URL.Path, elapsedMicros)
}

func writeCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Method", "GET,POST,DELETE,PUT,PATCH,OPTIONS,HEAD")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func ersp(msg string) (string, error) {
	return "", errors.New(msg)
}
