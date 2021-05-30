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
		return NewResult("", nil, req, act, act, &vaction.Folder{Req: req, Act: act}), nil
	case "all":
		return sourceAll(req, act)
	case "source", "package", "model":
		return sourceItem(req, act)
	case "static":
		return NewResult("", nil, req, act, act, &vaction.Static{Req: req, Act: act}), nil
	case "activity":
		return sourceActivity(req, act)
	default:
		page := &views.TODO{Message: "unhandled action type [" + act.Type + "]"}
		return NewResult("", nil, req, act, act, page), nil
	}
}

func sourceAll(req *cutil.WorkspaceRequest, act *action.Action) (*Result, error) {
	switch len(req.Path) {
	case 0:
		return NewResult("", nil, req, act, act, &vaction.Sources{Req: req, Act: act}), nil
	case 1:
		src := req.Sources.Get(req.Path[0])
		sch := req.Schemata[req.Path[0]]
		return NewResult("", nil, req, act, act, &vaction.Source{Req: req, Act: act, Source: src, Pkg: sch.ModelsByPackage()}), nil
	default:
		p, err := rootItemFor(req, req.Path[0])
		if err != nil {
			return ErrResult(req, act, err)
		}
		return process(req, act, p, req.Path[0], req.Path[1:])
	}
}

func sourceItem(req *cutil.WorkspaceRequest, act *action.Action) (*Result, error) {
	src := act.Config.GetStringOpt("source")
	p, err := rootItemFor(req, src)
	if err != nil {
		return ErrResult(req, act, err)
	}
	var x []string
	if act.Type == "package" || act.Type == "model" {
		t := act.Config.GetStringOpt(act.Type)
		if t == "" {
			return ErrResult(req, act, errors.New("must provide ["+act.Type+"] in config"))
		}
		x = util.SplitAndTrim(t, "/")
	}
	return process(req, act, p, src, append(x, req.Path...))
}

func process(req *cutil.WorkspaceRequest, act *action.Action, pkg *model.Package, srcKey string, path []string) (*Result, error) {
	i, remaining := pkg.Get(path)
	switch t := i.(type) {
	case *model.Model:
		return processModel(req, act, srcKey, t, remaining)
	case *model.Package:
		return processPackage(req, act, t)
	case error:
		return ErrResult(req, act, errors.Wrapf(t, "provided path [%v] can't be loaded", req.R.URL.Path))
	case nil:
		return ErrResult(req, act, errors.Errorf("nil path [%v] can't be loaded", req.R.URL.Path))
	default:
		return ErrResult(req, act, errors.Errorf("unhandled type: %T", t))
	}
}

func rootItemFor(req *cutil.WorkspaceRequest, srcKey string) (*model.Package, error) {
	if srcKey == "" {
		return nil, errors.New("must provide source key")
	}
	sch, ok := req.Schemata[srcKey]
	if !ok {
		return nil, errors.New("no schema registered for source [" + srcKey + "]")
	}
	return sch.ModelsByPackage(), nil
}
