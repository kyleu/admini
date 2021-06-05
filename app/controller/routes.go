package controller

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func BuildRouter() *router.Router {
	w := fasthttp.CompressHandler
	r := router.New()
	r.GET("/", w(Home))

	r.GET(defaultSearchPath, w(Search))
	r.GET(defaultProfilePath, w(Profile))
	r.GET("/settings", w(Settings))
	r.GET("/admin", w(Admin))

	r.GET("/help", w(Help))
	r.GET("/feedback", w(Feedback))

	r.GET("/source", w(SourceList))
	r.POST("/source", w(SourceInsert))
	r.GET("/source/_new", w(SourceNew))
	r.GET("/source/{key}", w(SourceDetail))
	r.GET("/source/{key}/edit", w(SourceEdit))
	r.POST("/source/{key}", w(SourceSave))
	r.GET("/source/{key}/refresh", w(SourceRefresh))
	r.GET("/source/{key}/delete", w(SourceDelete))

	r.GET("/project", w(ProjectList))
	r.POST("/project", w(ProjectInsert))
	r.GET("/project/_new", w(ProjectNew))
	r.GET("/project/{key}", w(ProjectDetail))
	r.POST("/project/{key}", w(ProjectSave))
	r.GET("/project/{key}/edit", w(ProjectEdit))
	r.POST("/project/{key}/actions", w(ActionOrdering))
	r.GET("/project/{key}/action/{_:*}", w(ActionEdit))
	r.POST("/project/{key}/action/{_:*}", w(ActionSave))
	r.GET("/project/{key}/test", w(ProjectTest))
	r.GET("/project/{key}/delete", w(ProjectDelete))

	r.GET("/x/{key}", w(WorkspaceProject))
	r.GET("/x/{key}/{_:*}", w(WorkspaceProject))
	r.GET("/s/{key}", w(WorkspaceSource))
	r.GET("/s/{key}/{_:*}", w(WorkspaceSource))

	r.GET("/sandbox", w(SandboxList))
	r.GET("/sandbox/{key}", w(SandboxRun))

	r.GET("/modules", w(Modules))

	r.GET("/favicon.ico", Favicon)
	r.GET("/robots.txt", RobotsTxt)
	r.GET("/assets/{_:*}", Static)

	r.OPTIONS("/", w(Options))
	r.OPTIONS("/{_:*}", w(Options))
	r.NotFound = NotFound

	return r
}
