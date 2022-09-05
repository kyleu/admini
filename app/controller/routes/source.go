package routes

import (
	"github.com/fasthttp/router"

	"admini.dev/admini/app/controller/csource"
)

func sourceRoutes(r *router.Router) {
	r.GET("/source", csource.SourceList)
	r.POST("/source", csource.SourceInsert)
	r.GET("/source/_new", csource.SourceNew)
	r.GET("/source/_example", csource.SourceExample)
	r.GET("/source/{key}", csource.SourceDetail)
	r.GET("/source/{key}/edit", csource.SourceEdit)
	r.POST("/source/{key}", csource.SourceSave)
	r.GET("/source/{key}/refresh", csource.SourceRefresh)
	r.GET("/source/{key}/delete", csource.SourceDelete)
	r.GET("/source/{key}/hack", csource.SourceHack)
	r.GET("/source/{key}/model/{path:*}", csource.SourceModelDetail)
	r.POST("/source/{key}/model/{path:*}", csource.SourceModelSave)
}
