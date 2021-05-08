package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/views"
)

const (
	workspaceProjectRoute = "workspace"
	workspaceSourceRoute  = "workspace.source"
)

type workspaceRequest struct {
	T    string
	K    string
	W    http.ResponseWriter
	R    *http.Request
	AS   *app.State
	PS   *cutil.PageState
	I    interface{}
	Path []string
	Src  *source.Source
}

func handle(req *workspaceRequest) (string, error) {
	switch t := req.I.(type) {
	case *model.Model:
		return handleModel(req, t)
	case *model.Package:
		return handlePackage(req, t)
	case error:
		return ersp("provided path [%v] can't be loaded: %+v", req.R.URL.Path, t)
	case nil:
		return ersp("nil path [%v] can't be loaded: %+v", req.R.URL.Path, t)
	default:
		return ersp("unhandled type: %T", t)
	}
}

func whoops(req *workspaceRequest, msg string) (string, error) {
	page := &views.TODO{Message: fmt.Sprintf("%v [%v]", msg, strings.Join(req.Path, "/"))}
	return render(req.R, req.W, req.AS, page, req.PS, req.Path...)
}
