// Package routes - Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"admini.dev/admini/app/controller"
	"admini.dev/admini/app/controller/clib"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/util"
)

func SiteRoutes(logger util.Logger) (http.Handler, error) {
	r := mux.NewRouter()

	makeRoute(r, http.MethodGet, cutil.DefaultProfilePath, clib.ProfileSite)
	makeRoute(r, http.MethodPost, cutil.DefaultProfilePath, clib.ProfileSave)
	makeRoute(r, http.MethodGet, "/auth/{key}", clib.AuthDetail)
	makeRoute(r, http.MethodGet, "/auth/callback/{key}", clib.AuthCallback)
	makeRoute(r, http.MethodGet, "/auth/logout/{key}", clib.AuthLogout)

	makeRoute(r, http.MethodGet, "/favicon.ico", clib.Favicon)
	makeRoute(r, http.MethodGet, "/assets/{path:.*}", clib.Static)

	makeRoute(r, http.MethodGet, "/", controller.Site)
	makeRoute(r, http.MethodGet, "/{path:.*}", controller.Site)

	makeRoute(r, http.MethodOptions, "/", controller.Options)
	r.HandleFunc("/", controller.NotFoundAction)

	return cutil.WireRouter(r, logger)
}
