// Package routes - Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller"
	"admini.dev/admini/app/controller/clib"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/util"
)

func makeRoute(x *mux.Router, method string, path string, f http.HandlerFunc) {
	cutil.AddRoute(method, path)
	x.HandleFunc(path, f).Methods(method)
}

//nolint:revive
func AppRoutes(as *app.State, logger util.Logger) (http.Handler, error) {
	r := mux.NewRouter()

	makeRoute(r, http.MethodGet, "/", controller.Home)
	makeRoute(r, http.MethodGet, "/healthcheck", clib.Healthcheck)
	makeRoute(r, http.MethodGet, "/about", clib.About)

	makeRoute(r, http.MethodGet, cutil.DefaultProfilePath, clib.Profile)
	makeRoute(r, http.MethodPost, cutil.DefaultProfilePath, clib.ProfileSave)
	makeRoute(r, http.MethodGet, "/auth/{key}", clib.AuthDetail)
	makeRoute(r, http.MethodGet, "/auth/callback/{key}", clib.AuthCallback)
	makeRoute(r, http.MethodGet, "/auth/logout/{key}", clib.AuthLogout)
	makeRoute(r, http.MethodGet, cutil.DefaultSearchPath, clib.Search)

	themeRoutes(r)

	// $PF_SECTION_START(routes)$
	makeRoute(r, http.MethodGet, "/refresh", controller.Refresh)

	sourceRoutes(r)
	projectRoutes(r)

	makeRoute(r, http.MethodGet, "/x/{key}", controller.WorkspaceProject)
	r.PathPrefix("/x/{key}/").Methods(http.MethodGet, http.MethodPost).HandlerFunc(controller.WorkspaceProject)

	makeRoute(r, http.MethodGet, "/s/{key}", controller.WorkspaceSource)
	r.PathPrefix("/s/{key}/").Methods(http.MethodGet, http.MethodPost).HandlerFunc(controller.WorkspaceSource)
	// $PF_SECTION_END(routes)$

	makeRoute(r, http.MethodGet, "/admin", clib.Admin)
	makeRoute(r, http.MethodGet, "/admin/", clib.Admin)
	makeRoute(r, http.MethodPost, "/admin/", clib.Admin)
	makeRoute(r, http.MethodGet, "/admin/sandbox", controller.SandboxList)
	makeRoute(r, http.MethodGet, "/admin/sandbox/{key}", controller.SandboxRun)

	makeRoute(r, http.MethodGet, "/favicon.ico", clib.Favicon)
	makeRoute(r, http.MethodGet, "/robots.txt", clib.RobotsTxt)
	makeRoute(r, http.MethodGet, "/assets/{path:.*}", clib.Static)

	makeRoute(r, http.MethodOptions, "/", controller.Options)
	r.HandleFunc("/", controller.NotFoundAction)

	return cutil.WireRouter(r, logger)
}
