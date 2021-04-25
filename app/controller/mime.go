package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kyleu/admini/app/util"
)

func respondJSON(w http.ResponseWriter, filename string, body interface{}) (string, error) {
	return respondMIME(filename, "application/json", "json", util.ToJSONBytes(body, true), w)
}

func respondMIME(filename string, mime string, ext string, ba []byte, w http.ResponseWriter) (string, error) {
	w.Header().Set("Content-Type", mime+"; charset=UTF-8")
	if len(filename) > 0 {
		if !strings.HasSuffix(filename, "."+ext) {
			filename = filename + "." + ext
		}
		w.Header().Set("Content-Disposition", `attachment; filename="`+filename+`"`)
	}
	writeCORS(w)
	if len(ba) == 0 {
		return "", fmt.Errorf("no bytes available to write")
	}
	if _, err := w.Write(ba); err != nil {
		return "", fmt.Errorf("cannot write to response: %w", err)
	}

	return "", nil
}

func writeCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Method", "GET,POST,DELETE,PUT,PATCH,OPTIONS,HEAD")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func getContentType(r *http.Request) string {
	ret := r.Header.Get("Content-Type")
	if idx := strings.Index(ret, ";"); idx > -1 {
		ret = ret[0:idx]
	}
	return strings.TrimSpace(ret)
}

func isContentTypeJSON(c string) bool {
	return c == "application/json" || c == "text/json"
}
