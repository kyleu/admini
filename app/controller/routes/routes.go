// Package routes - Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller"
	"admini.dev/admini/app/controller/clib"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/util"
)

//nolint:revive
func AppRoutes(as *app.State, logger util.Logger) fasthttp.RequestHandler {
	r := router.New()

	r.GET("/", controller.Home)
	r.GET("/healthcheck", clib.Healthcheck)
	r.GET("/about", clib.About)

	r.GET(cutil.DefaultProfilePath, clib.Profile)
	r.POST(cutil.DefaultProfilePath, clib.ProfileSave)
	r.GET("/auth/{key}", clib.AuthDetail)
	r.GET("/auth/callback/{key}", clib.AuthCallback)
	r.GET("/auth/logout/{key}", clib.AuthLogout)
	r.GET(cutil.DefaultSearchPath, clib.Search)
	themeRoutes(r)

	// $PF_SECTION_START(routes)$
	r.GET("/refresh", controller.Refresh)

	sourceRoutes(r)
	projectRoutes(r)

	r.GET("/x/{key}", controller.WorkspaceProject)
	r.GET("/x/{key}/{_:*}", controller.WorkspaceProject)
	r.POST("/x/{key}/{_:*}", controller.WorkspaceProject)
	r.GET("/s/{key}", controller.WorkspaceSource)
	r.GET("/s/{key}/{_:*}", controller.WorkspaceSource)
	r.POST("/s/{key}/{_:*}", controller.WorkspaceSource)
	// $PF_SECTION_END(routes)$

	r.GET("/admin", clib.Admin)
	r.GET("/admin/sandbox", controller.SandboxList)
	r.GET("/admin/sandbox/{key}", controller.SandboxRun)
	r.GET("/admin/{path:*}", clib.Admin)
	r.POST("/admin/{path:*}", clib.Admin)

	r.GET("/favicon.ico", clib.Favicon)
	r.GET("/robots.txt", clib.RobotsTxt)
	r.GET("/assets/{_:*}", clib.Static)

	r.OPTIONS("/", controller.Options)
	r.OPTIONS("/{_:*}", controller.Options)
	r.NotFound = controller.NotFoundAction

	return clib.WireRouter(r, logger)
}
