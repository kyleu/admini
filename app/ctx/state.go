package ctx

import "github.com/gorilla/mux"

var ActiveRouter *mux.Router

type PageState struct {
	Breadcrumbs []string
	Router      *mux.Router
}
