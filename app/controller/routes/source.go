package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"admini.dev/admini/app/controller/csource"
)

func sourceRoutes(r *mux.Router) {
	makeRoute(r, http.MethodGet, "/source", csource.SourceList)
	makeRoute(r, http.MethodPost, "/source", csource.SourceInsert)
	makeRoute(r, http.MethodGet, "/source/_new", csource.SourceNew)
	makeRoute(r, http.MethodGet, "/source/_example", csource.SourceExample)
	makeRoute(r, http.MethodGet, "/source/{key}", csource.SourceDetail)
	makeRoute(r, http.MethodGet, "/source/{key}/edit", csource.SourceEdit)
	makeRoute(r, http.MethodPost, "/source/{key}", csource.SourceSave)
	makeRoute(r, http.MethodGet, "/source/{key}/refresh", csource.SourceRefresh)
	makeRoute(r, http.MethodGet, "/source/{key}/delete", csource.SourceDelete)
	makeRoute(r, http.MethodGet, "/source/{key}/hack", csource.SourceHack)
	makeRoute(r, http.MethodGet, "/source/{key}/model/{path:.*}", csource.SourceModelDetail)
	makeRoute(r, http.MethodPost, "/source/{key}/model/{path:.*}", csource.SourceModelSave)
}
