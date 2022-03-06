// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"admini.dev/app/lib/telemetry/httpmetrics"
	"admini.dev/app/util"
)

//nolint
func AppRoutes() fasthttp.RequestHandler {
	r := router.New()

	r.GET("/", Home)
	r.GET("/healthcheck", Healthcheck)
	r.GET("/about", About)
	r.GET("/theme", ThemeList)
	r.GET("/theme/{key}", ThemeEdit)
	r.POST("/theme/{key}", ThemeSave)
	r.GET(defaultSearchPath, Search)

	r.GET(defaultProfilePath, Profile)
	r.POST(defaultProfilePath, ProfileSave)
	r.GET("/auth/{key}", AuthDetail)
	r.GET("/auth/callback/{key}", AuthCallback)
	r.GET("/auth/logout/{key}", AuthLogout)

	// $PF_SECTION_START(routes)$
	r.GET("/refresh", Refresh)

	r.GET("/source", SourceList)
	r.POST("/source", SourceInsert)
	r.GET("/source/_new", SourceNew)
	r.GET("/source/_example", SourceExample)
	r.GET("/source/{key}", SourceDetail)
	r.GET("/source/{key}/edit", SourceEdit)
	r.POST("/source/{key}", SourceSave)
	r.GET("/source/{key}/refresh", SourceRefresh)
	r.GET("/source/{key}/delete", SourceDelete)
	r.GET("/source/{key}/hack", SourceHack)
	r.GET("/source/{key}/model/{path:*}", SourceModelDetail)
	r.POST("/source/{key}/model/{path:*}", SourceModelSave)

	r.GET("/project", ProjectList)
	r.POST("/project", ProjectInsert)
	r.GET("/project/_new", ProjectNew)
	r.GET("/project/{key}", ProjectDetail)
	r.POST("/project/{key}", ProjectSave)
	r.GET("/project/{key}/edit", ProjectEdit)
	r.POST("/project/{key}/actions", ActionOrdering)
	r.GET("/project/{key}/action/{path:*}", ActionEdit)
	r.POST("/project/{key}/action/{path:*}", ActionSave)
	r.GET("/project/{key}/test", ProjectTest)
	r.GET("/project/{key}/delete", ProjectDelete)

	r.GET("/x/{key}", WorkspaceProject)
	r.GET("/x/{key}/{_:*}", WorkspaceProject)
	r.POST("/x/{key}/{_:*}", WorkspaceProject)
	r.GET("/s/{key}", WorkspaceSource)
	r.GET("/s/{key}/{_:*}", WorkspaceSource)
	r.POST("/s/{key}/{_:*}", WorkspaceSource)
	// $PF_SECTION_END(routes)$

	r.GET("/admin", Admin)
	r.GET("/admin/sandbox", SandboxList)
	r.GET("/admin/sandbox/{key}", SandboxRun)
	r.GET("/admin/{path:*}", Admin)

	r.GET("/favicon.ico", Favicon)
	r.GET("/robots.txt", RobotsTxt)
	r.GET("/assets/{_:*}", Static)

	r.OPTIONS("/", Options)
	r.OPTIONS("/{_:*}", Options)
	r.NotFound = NotFound

	p := httpmetrics.NewMetrics(util.AppKey)
	return fasthttp.CompressHandlerBrotliLevel(p.WrapHandler(r), fasthttp.CompressBrotliBestSpeed, fasthttp.CompressBestSpeed)
}
