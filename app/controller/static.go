package controller

import (
	"net/http"
	"path/filepath"
	"runtime/debug"
	"strings"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/views/vhelp"
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/assets"
)

const assetBase = "assets"

func Favicon(w http.ResponseWriter, r *http.Request) {
	data, hash, contentType, err := assets.Asset(assetBase, "/favicon.ico")
	ZipResponse(w, r, data, hash, contentType, err)
}

func RobotsTxt(w http.ResponseWriter, r *http.Request) {
	data, hash, contentType, err := assets.Asset(assetBase, "/robots.txt")
	ZipResponse(w, r, data, hash, contentType, err)
}

func Modules(w http.ResponseWriter, r *http.Request) {
	act("modules", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		mods, ok := debug.ReadBuildInfo()
		if !ok {
			return "", errors.New("unable to gather modules")
		}
		ps.Title = "Modules"
		ps.Data = mods.Deps
		return render(r, w, as, &vhelp.Modules{Mods: mods.Deps}, ps, "modules")
	})
}

func Routes(w http.ResponseWriter, r *http.Request) {
	act("routes", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		routes := cutil.ExtractRoutes(as.Router)
		ps.Title = "Routes"
		ps.Data = routes
		return render(r, w, as, &vhelp.Routes{Routes: routes}, ps, "routes")
	})
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
