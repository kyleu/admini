package controller

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/kyleu/admini/app/assets"
)

const assetBase = "web/assets"

func Favicon(w http.ResponseWriter, r *http.Request) {
	data, hash, contentType, err := assets.Asset(assetBase, "/favicon.ico")
	ZipResponse(w, r, data, hash, contentType, err)
}

func RobotsTxt(w http.ResponseWriter, r *http.Request) {
	data, hash, contentType, err := assets.Asset(assetBase, "/robots.txt")
	ZipResponse(w, r, data, hash, contentType, err)
}

func Static(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(strings.TrimPrefix(r.URL.Path, "/assets"))
	if err == nil {
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		data, hash, contentType, e := assets.Asset(assetBase, path)
		ZipResponse(w, r, data, hash, contentType, e)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func ZipResponse(w http.ResponseWriter, r *http.Request, data []byte, hash string, contentType string, err error) {
	if err == nil {
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Type", contentType)
		// w.Header().Add("Cache-Control", "public, max-age=31536000")
		w.Header().Add("ETag", hash)
		if r.Header.Get("If-None-Match") == hash {
			w.WriteHeader(http.StatusNotModified)
		} else {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(data)
		}
	} else {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
