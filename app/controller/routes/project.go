package routes

import (
	"admini.dev/admini/app/controller"
	"admini.dev/admini/app/controller/cproject"
	"github.com/fasthttp/router"
)

func projectRoutes(r *router.Router) {
	r.GET("/project", cproject.ProjectList)
	r.POST("/project", cproject.ProjectInsert)
	r.GET("/project/_new", cproject.ProjectNew)
	r.GET("/project/{key}", cproject.ProjectDetail)
	r.POST("/project/{key}", cproject.ProjectSave)
	r.GET("/project/{key}/edit", cproject.ProjectEdit)
	r.POST("/project/{key}/actions", controller.ActionOrdering)
	r.GET("/project/{key}/action/{path:*}", controller.ActionEdit)
	r.POST("/project/{key}/action/{path:*}", controller.ActionSave)
	r.GET("/project/{key}/test", cproject.ProjectTest)
	r.GET("/project/{key}/delete", cproject.ProjectDelete)
}
