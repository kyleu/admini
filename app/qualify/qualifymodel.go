package qualify

import (
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
)

func qualifyModel(req *Request, act *action.Action, srcKey string, modelPath util.Pkg, schemata schema.Schemata) (Results, error) {
	if act.Type == action.TypeAll {
		sch, err := schemata.GetWithError(srcKey)
		if err != nil {
			return nil, errors.Wrapf(err, "can't find source in project with key [%s]", modelPath[0])
		}
		src := sch.ModelsByPackage()
		m, _ := src.Get(modelPath)
		if m != nil {
			return modelResults(req, act.Path(), act.IconWithFallback(), append([]string{srcKey}, modelPath...), "all")
		}
		return nil, nil
	}

	if !checkSource(act, srcKey) {
		return nil, nil
	}
	switch act.Type {
	case action.TypeSource:
		src, err := action.GetSource(act, schemata)
		if err != nil {
			return nil, err
		}
		m, _ := src.Get(modelPath)
		if m != nil {
			return modelResults(req, act.Path(), act.IconWithFallback(), modelPath, "src")
		}
	case action.TypePackage:
		pkg, remaining, err := action.GetPackage(act, schemata)
		if err != nil {
			return nil, err
		}
		if modelPath.StartsWith(pkg.Path()) {
			m, _ := pkg.Get(remaining)
			if m != nil {
				return modelResults(req, act.Path(), act.IconWithFallback(), modelPath[len(pkg.Path()):], "pkg")
			}
		}
	case action.TypeModel:
		model, remaining, err := action.GetModel(act, schemata)
		if err != nil {
			return nil, err
		}

		if modelPath.StartsWith(model.Path()) {
			return modelResults(req, act.Path(), act.IconWithFallback(), remaining, "model")
		}
	}
	return nil, nil
}

func modelResults(req *Request, actPath []string, icon string, remaining []string, dbg string) (Results, error) {
	path := append([]string{}, remaining...)
	if req.Action == "view" {
		keys, err := req.Params.GetStringArray("keys", false)
		if err != nil {
			return nil, err
		}
		path = append(append(path, "v"), keys...)
	}
	return Results{&Result{Action: actPath, Icon: icon, Path: path, Debug: dbg}}, nil
}

func checkSource(act *action.Action, key string) bool {
	k, err := act.Config.GetString("source", false)
	if err != nil {
		return false
	}
	return k == key
}
