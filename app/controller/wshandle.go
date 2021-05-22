package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/project/action"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/views"
)

const (
	pRoute = "workspace"
	sRoute = "workspace.source"
)

type workspaceRequest struct {
	T       string
	K       string
	W       http.ResponseWriter
	R       *http.Request
	AS      *app.State
	PS      *cutil.PageState
	Source  string
	Item    interface{}
	Path    []string
	Project *project.Project
	Sources source.Sources
}

func handle(req *workspaceRequest) (string, error) {
	switch t := req.Item.(type) {
	case *action.Action:
		return handleAction(req, t)
	case *model.Model:
		return handleModel(req, t)
	case *model.Package:
		return handlePackage(req, t)
	case error:
		return "", errors.Wrap(t, fmt.Sprintf("provided path [%v] can't be loaded", req.R.URL.Path))
	case nil:
		return ersp("nil path [%v] can't be loaded", req.R.URL.Path)
	default:
		return ersp("unhandled type: %T", t)
	}
}

func whoops(req *workspaceRequest, msg string, path ...string) (string, error) {
	page := &views.TODO{Message: fmt.Sprintf("%v [%v]", msg, strings.Join(req.Path, "/"))}
	return render(req.R, req.W, req.AS, page, req.PS, path...)
}
