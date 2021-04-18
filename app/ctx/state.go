package ctx

import (
	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/menu"
	"github.com/kyleu/admini/app/source"
)

type AppState struct {
	Router  *mux.Router
	Files   filesystem.FileLoader
	Sources *source.Service
}

type PageState struct {
	Menu        menu.Items
	Breadcrumbs []string
}
