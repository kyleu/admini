package workspace

import (
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/project/action"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/views"
	"github.com/kyleu/admini/views/layout"
	"github.com/kyleu/admini/views/vaction"
	"github.com/kyleu/admini/views/vworkspace"
	"github.com/pkg/errors"
	"strings"
)

type Result struct {
	Title       string      `json:"title"`
	Breadcrumbs []string    `json:"bc"`
	Data        interface{} `json:"data"`
	Page        layout.Page `json:"page"`
}

func NewResult(title string, bc []string, req *cutil.WorkspaceRequest, act *action.Action, data interface{}, page layout.Page) *Result {
	if title == "" && act != nil {
		title = act.Name()
	}
	if bc == nil && act != nil {
		bc = append(act.Path(), req.Path...)
		println("@@@ " + strings.Join(act.Path(), "//"))
		println("### " + strings.Join(req.Path, "//"))
		println("::: " + strings.Join(bc, "//"))
	}
	return &Result{Title: title, Breadcrumbs: bc, Data: data, Page: page}
}

func ErrResult(req *cutil.WorkspaceRequest, act *action.Action, err error) (*Result, error) {
	title := "Error"
	if act != nil {
		title += ": " + act.Name()
	}
	return NewResult(title, nil, req, act, "ERROR: "+err.Error(), nil), err
}

func ActionHandler(req *cutil.WorkspaceRequest, act *action.Action) (*Result, error) {
	if act == nil || (act.Key == "" && len(act.Pkg) == 0) {
		return NewResult(req.Project.Name(), nil, req, act, req.Project, &vworkspace.WorkspaceOverview{Req: req}), nil
	}
	switch act.Type {
	case "", "folder":
		return NewResult("", nil, req, act, act, &vaction.ResultFolder{Req: req, Act: act}), nil
	case "source", "package", "model":
		return sourceItem(req, act)
	case "static":
		return NewResult("", nil, req, act, act, &vaction.ResultStatic{Req: req, Act: act}), nil
	case "activity":
		return sourceActivity(req, act)
	default:
		page := &views.TODO{Message: "unhandled action type [" + act.Type + "]"}
		return NewResult("", nil, req, act, act, page), nil
	}
}

func sourceItem(req *cutil.WorkspaceRequest, act *action.Action) (*Result, error) {
	p, err := itemFor(req, act)
	if err != nil {
		return ErrResult(req, act, err)
	}
	var x []string
	if act.Type == "package" || act.Type == "model" {
		t := act.Config[act.Type]
		if t == "" {
			return ErrResult(req, act, errors.New("must provide ["+act.Type+"] in config"))
		}
		x = util.SplitAndTrim(t, "/")
	}
	return process(req, act, p, append(x, req.Path...))
}

func process(req *cutil.WorkspaceRequest, act *action.Action, pkg *model.Package, path []string) (*Result, error) {
	i, remaining := pkg.Get(path)
	switch t := i.(type) {
	case *model.Model:
		return processModel(req, act, pkg, path, t, remaining)
	case *model.Package:
		return processPackage(req, act, pkg, path, t, remaining)
	case error:
		return ErrResult(req, act, errors.Wrapf(t, "provided path [%v] can't be loaded", req.R.URL.Path))
	case nil:
		return ErrResult(req, act, errors.Errorf("nil path [%v] can't be loaded", req.R.URL.Path))
	default:
		return ErrResult(req, act, errors.Errorf("unhandled type: %T", t))
	}
}

func itemFor(req *cutil.WorkspaceRequest, act *action.Action) (*model.Package, error) {
	srcKey := act.Config["source"]
	if srcKey == "" {
		return nil, errors.New("must provide source key")
	}
	sch, ok := req.Schemata[srcKey]
	if !ok {
		return nil, errors.New("no schema registered for source [" + srcKey + "]")
	}
	return sch.ModelsByPackage(), nil
}
