package workspace

import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/action"
	model2 "github.com/kyleu/admini/app/schema/model"
	"github.com/kyleu/admini/app/telemetry"
	"go.opentelemetry.io/otel/attribute"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/views/vaction"
	"github.com/kyleu/admini/views/vworkspace"
	"github.com/pkg/errors"
)

func ActionHandler(req *cutil.WorkspaceRequest, act *action.Action, as *app.State) (*Result, error) {
	ctx, span := telemetry.StartSpan(req.Context, "workspace", act.String())
	req.Context = ctx
	defer span.End()

	if act == nil || (act.Key == "" && len(act.Pkg) == 0) {
		return NewResult(req.Project.Name(), nil, req, act, req.Project, &vworkspace.WorkspaceOverview{Req: req}), nil
	}
	span.SetAttributes(attribute.String("actType", act.Type.Key), attribute.String("actPkg", act.Pkg.String()), attribute.String("actKey", act.Key))
	switch act.Type {
	case action.TypeFolder:
		return NewResult("", nil, req, act, act, &vaction.Folder{Req: req, Act: act}), nil
	case action.TypeStatic:
		return NewResult("", nil, req, act, act, &vaction.Static{Req: req, Act: act}), nil

	case action.TypeAll:
		return sourceAll(req, act, as)
	case action.TypeSource, action.TypePackage, action.TypeModel:
		return sourceItem(req, act, as)
	case action.TypeActivity:
		return sourceActivity(req, act, as)

	default:
		return nil, errors.Errorf("unhandled action type [%s]", act.Type.Key)
	}
}

func sourceAll(req *cutil.WorkspaceRequest, act *action.Action, as *app.State) (*Result, error) {
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
		return process(req, act, p, req.Path[0], req.Path[1:], as)
	}
}

func sourceItem(req *cutil.WorkspaceRequest, act *action.Action, as *app.State) (*Result, error) {
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
		x = util.StringSplitAndTrim(t, "/")
	}
	return process(req, act, p, src, append(x, req.Path...), as)
}

func process(req *cutil.WorkspaceRequest, act *action.Action, pkg *model2.Package, srcKey string, path []string, as *app.State) (*Result, error) {
	i, remaining := pkg.Get(path)
	switch t := i.(type) {
	case *model2.Model:
		return processModel(req, act, srcKey, t, remaining, as)
	case *model2.Package:
		return processPackage(req, act, t)
	case error:
		return ErrResult(req, act, errors.Wrapf(t, "provided path [%s] can't be loaded", string(req.Ctx.URI().Path())))
	case nil:
		return ErrResult(req, act, errors.Errorf("nil path [%s] can't be loaded", string(req.Ctx.URI().Path())))
	default:
		return ErrResult(req, act, errors.Errorf("unhandled type for processing: %T", t))
	}
}

func rootItemFor(req *cutil.WorkspaceRequest, srcKey string) (*model2.Package, error) {
	if srcKey == "" {
		return nil, errors.New("must provide source key")
	}
	sch, ok := req.Schemata[srcKey]
	if !ok {
		return nil, errors.Errorf("no schema registered for source [%s]", srcKey)
	}
	return sch.ModelsByPackage(), nil
}
