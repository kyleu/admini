package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"admini.dev/admini/app/controller"
	"admini.dev/admini/app/controller/cproject"
)

func projectRoutes(r *mux.Router) {
	makeRoute(r, http.MethodGet, "/project", cproject.ProjectList)
	makeRoute(r, http.MethodPost, "/project", cproject.ProjectInsert)
	makeRoute(r, http.MethodGet, "/project/_new", cproject.ProjectNew)
	makeRoute(r, http.MethodGet, "/project/{key}", cproject.ProjectDetail)
	makeRoute(r, http.MethodPost, "/project/{key}", cproject.ProjectSave)
	makeRoute(r, http.MethodGet, "/project/{key}/edit", cproject.ProjectEdit)
	makeRoute(r, http.MethodPost, "/project/{key}/actions", controller.ActionOrdering)
	makeRoute(r, http.MethodGet, "/project/{key}/action/{path:.*}", controller.ActionEdit)
	makeRoute(r, http.MethodPost, "/project/{key}/action/{path:.*}", controller.ActionSave)
	makeRoute(r, http.MethodGet, "/project/{key}/test", cproject.ProjectTest)
	makeRoute(r, http.MethodGet, "/project/{key}/delete", cproject.ProjectDelete)
}
