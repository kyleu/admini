package controller

import (
	"github.com/fasthttp/router"
)

func BuildRouterNew() *router.Router {
	r := router.New()
	r.GET("/", Home)

	r.GET(defaultSearchPath, Search)
	r.GET(defaultProfilePath, Profile)
	r.GET("/settings", Settings)
	r.GET("/admin", Admin)

	r.GET("/help", Help)
	r.GET("/feedback", Feedback)

	r.GET("/source", SourceList)
	r.POST("/source", SourceInsert)
	r.GET("/source/_new", SourceNew)
	r.GET("/source/{key}", SourceDetail)
	r.GET("/source/{key}/edit", SourceEdit)
	r.POST("/source/{key}", SourceSave)
	r.GET("/source/{key}/refresh", SourceRefresh)
	r.GET("/source/{key}/delete", SourceDelete)

	r.GET("/project", ProjectList)
	r.POST("/project", ProjectInsert)
	r.GET("/project/_new", ProjectNew)
	r.GET("/project/{key}", ProjectDetail)
	r.POST("/project/{key}", ProjectSave)
	r.GET("/project/{key}/edit", ProjectEdit)
	r.POST("/project/{key}/actions", ActionOrdering)
	r.GET("/project/{key}/action/{_:*}", ActionEdit)
	r.POST("/project/{key}/action/{_:*}", ActionSave)
	r.GET("/project/{key}/test", ProjectTest)
	r.GET("/project/{key}/delete", ProjectDelete)

	r.GET("/x/{key}", WorkspaceProject)
	r.GET("/x/{key}/{_:*}", WorkspaceProject)
	r.GET("/s/{key}", WorkspaceSource)
	r.GET("/s/{key}/{_:*}", WorkspaceSource)

	r.GET("/sandbox", SandboxList)
	r.GET("/sandbox/{key}", SandboxRun)

	r.GET("/modules", Modules)

	r.GET("/favicon.ico", Favicon)
	r.GET("/robots.txt", RobotsTxt)
	r.GET("/assets/{_:*}", Static)

	r.OPTIONS("/", Options)
	r.OPTIONS("/{_:*}", Options)
	r.NotFound = NotFound

	return r
}
