package workspace

import (
	"github.com/pkg/errors"

	"admini.dev/admini/app"
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/lib/telemetry"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/vaction"
	"admini.dev/admini/views/vworkspace"
)

func ActionHandler(req *cutil.WorkspaceRequest, act *action.Action, as *app.State) (*Result, error) {
	ctx, span, _ := telemetry.StartSpan(req.Context, "workspace:"+act.String(), req.PS.Logger)
	req.Context = ctx
	defer span.Complete()

	if act == nil || (act.Key == "" && len(act.Pkg) == 0) {
		return NewResult(req.Project.Name(), nil, req, act, req.Project, &vworkspace.WorkspaceOverview{Req: req}), nil
	}
	span.Attributes(
		&telemetry.Attribute{Key: "actType", Value: act.TypeKey},
		&telemetry.Attribute{Key: "actPkg", Value: act.Pkg.String()},
		&telemetry.Attribute{Key: "actKey", Value: act.Key},
	)
	switch act.TypeKey {
	case action.TypeFolder.Key:
		return NewResult("", nil, req, act, act, &vaction.Folder{Req: req, Act: act}), nil
	case action.TypeStatic.Key:
		return NewResult("", nil, req, act, act, &vaction.Static{Req: req, Act: act}), nil

	case action.TypeAll.Key:
		return sourceAll(req, act, as)
	case action.TypeSource.Key, action.TypePackage.Key, action.TypeModel.Key:
		return sourceItem(req, act, as)
	case action.TypeActivity.Key:
		return sourceActivity(req, act, as)

	default:
		return nil, errors.Errorf("unhandled action type [%s]", act.TypeKey)
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
	switch act.TypeKey {
	case action.TypePackage.Key, action.TypeModel.Key:
		t := act.Config.GetStringOpt(act.TypeKey)
		if t == "" {
			return ErrResult(req, act, errors.Errorf("must provide [%s] in config", act.TypeKey))
		}
		x = util.StringSplitAndTrim(t, "/")
	}
	return process(req, act, p, src, append(x, req.Path...), as)
}

func process(req *cutil.WorkspaceRequest, act *action.Action, pkg *model.Package, srcKey string, path []string, as *app.State) (*Result, error) {
	i, remaining := pkg.Get(path)
	switch t := i.(type) {
	case *model.Model:
		return processModel(req, act, srcKey, t, remaining, as)
	case *model.Package:
		return processPackage(req, act, t)
	case error:
		return ErrResult(req, act, errors.Wrapf(t, "provided path [%s] can't be loaded", string(req.RC.URI().Path())))
	case nil:
		return ErrResult(req, act, errors.Errorf("nil path [%s] can't be loaded", string(req.RC.URI().Path())))
	default:
		return ErrResult(req, act, errors.Errorf("unhandled type for processing: %T", t))
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
