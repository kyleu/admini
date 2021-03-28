package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/kyleu/admini/app/util"
)

func act(key string, w http.ResponseWriter, r *http.Request, f func() (string, error)) {
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
}

func writeCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Method", "GET,POST,DELETE,PUT,PATCH,OPTIONS,HEAD")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func isContentTypeJSON(c string) bool {
	return c == "application/json" || c == "text/json"
}

func getContentType(r *http.Request) string {
	ret := r.Header.Get("Content-Type")
	idx := strings.Index(ret, ";")
	if idx > -1 {
		ret = ret[0:idx]
	}
	return strings.TrimSpace(ret)
}

func respondJSON(w http.ResponseWriter, filename string, body interface{}) (string, error) {
	b := util.ToJSONBytes(body, true)
	return respondMIME(filename, "application/json", "json", b, w)
}

func respondMIME(filename string, mime string, ext string, ba []byte, w http.ResponseWriter) (string, error) {
	w.Header().Set("Content-Type", mime+"; charset=UTF-8")
	if len(filename) > 0 {
		if !strings.HasSuffix(filename, "."+ext) {
			filename = filename + "." + ext
		}
		w.Header().Set("Content-Disposition", `attachment; filename="`+filename+`"`)
	}
	if len(ba) == 0 {
		return "", errors.New("no bytes available to write")
	}
	_, err := w.Write(ba)
	return "", fmt.Errorf("cannot write to response: %w", err)
}

func tmpl(_ int, err error) (string, error) {
	return "", err
}
