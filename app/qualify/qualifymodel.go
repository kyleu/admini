package qualify

import (
	"github.com/pkg/errors"

	"admini.dev/admini/app/action"
	"admini.dev/admini/app/lib/schema"
	"admini.dev/admini/app/util"
)

func qualifyModel(req *Request, act *action.Action, srcKey string, modelPkg util.Pkg, schemata schema.Schemata) (Qualifications, error) {
	modelPath := modelPkg
	if act.TypeKey == action.TypeAll.Key {
		sch, err := schemata.GetWithError(srcKey)
		if err != nil {
			return nil, errors.Wrapf(err, "can't find source in project with key [%s]", modelPath[0])
		}
		src := sch.ModelsByPackage()
		m, _ := src.Get(modelPath)
		if m != nil {
			return modelResults(req, act, act.IconWithFallback(), append([]string{srcKey}, modelPath...), "all")
		}
		return nil, nil
	}

	if !checkSource(act, srcKey) {
		return nil, nil
	}
	switch act.TypeKey {
	case action.TypeSource.Key:
		src, err := action.GetSource(act, schemata)
		if err != nil {
			return nil, err
		}
		m, _ := src.Get(modelPath)
		if m != nil {
			return modelResults(req, act, act.IconWithFallback(), modelPath, "src")
		}
	case action.TypePackage.Key:
		pkg, remaining, err := action.GetPackage(act, schemata)
		if err != nil {
			return nil, err
		}
		if modelPkg.StartsWith(pkg.Path()) {
			m, _ := pkg.Get(remaining)
			if m != nil {
				return modelResults(req, act, act.IconWithFallback(), modelPath[len(pkg.Path()):], "pkg")
			}
		}
	case action.TypeModel.Key:
		model, remaining, err := action.GetModel(act, schemata)
		if err != nil {
			return nil, err
		}

		if modelPkg.StartsWith(model.Path()) {
			return modelResults(req, act, act.IconWithFallback(), remaining, "model")
		}
	}
	return nil, nil
}

func modelResults(req *Request, act *action.Action, icon string, remaining []string, dbg string) (Qualifications, error) {
	path := append([]string{}, remaining...)
	if req.Action == "view" {
		keys, err := req.Params.GetStringArray("keys", false)
		if err != nil {
			return nil, err
		}
		path = append(append(path, "v"), keys...)
	}
	return Qualifications{&Qualification{Action: act, Icon: icon, Path: path, Debug: dbg}}, nil
}

func checkSource(act *action.Action, key string) bool {
	k, err := act.Config.GetString(action.TypeSource.Key, false)
	if err != nil {
		return false
	}
	return k == key
}
