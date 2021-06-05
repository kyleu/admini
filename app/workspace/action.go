package workspace

import (
	"fmt"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/project/action"
	"github.com/kyleu/admini/app/util"
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
		title = fmt.Sprintf("%s: %s", title, act.Name())
	}
	return NewResult(title, nil, req, act, fmt.Sprintf("ERROR: %s", err.Error()), nil), err
}

func ActionHandler(req *cutil.WorkspaceRequest, act *action.Action) (*Result, error) {
	if act == nil || (act.Key == "" && len(act.Pkg) == 0) {
		return NewResult(req.Project.Name(), nil, req, act, req.Project, &vworkspace.WorkspaceOverview{Req: req}), nil
	}
	switch act.Type {
	case action.TypeFolder:
		return NewResult("", nil, req, act, act, &vaction.Folder{Req: req, Act: act}), nil
	case action.TypeStatic:
		return NewResult("", nil, req, act, act, &vaction.Static{Req: req, Act: act}), nil

	case action.TypeAll:
		return sourceAll(req, act)
	case action.TypeSource, action.TypePackage, action.TypeModel:
		return sourceItem(req, act)
	case action.TypeActivity:
		return sourceActivity(req, act)

	default:
		return nil, errors.Errorf("unhandled action type [%s]", act.Type.Key)
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
	src := act.Config.GetStringOpt(action.TypeSource.Key)
	p, err := rootItemFor(req, src)
	if err != nil {
		return ErrResult(req, act, err)
	}
	var x []string
	if act.Type == action.TypePackage || act.Type == action.TypeModel {
		t := act.Config.GetStringOpt(act.Type.Key)
		if t == "" {
			return ErrResult(req, act, errors.Errorf("must provide [%s] in config", act.Type.Key))
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
		return ErrResult(req, act, errors.Wrapf(t, "provided path [%s] can't be loaded", string(req.Ctx.URI().Path())))
	case nil:
		return ErrResult(req, act, errors.Errorf("nil path [%s] can't be loaded", string(req.Ctx.URI().Path())))
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
		return nil, errors.Errorf("no schema registered for source [%s]", srcKey)
	}
	return sch.ModelsByPackage(), nil
}
