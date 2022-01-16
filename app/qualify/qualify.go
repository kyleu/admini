package qualify

import (
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/lib/schema"
	"github.com/kyleu/admini/app/util"
)

func Qualify(req *Request, acts action.Actions, schemata schema.Schemata) (Qualifications, error) {
	var ret Qualifications
	for _, act := range acts {
		childResult, err := qualifyAct(req, act, schemata)
		if err != nil {
			return nil, err
		}
		ret = append(ret, childResult...)
	}
	return ret, nil
}

func qualifyAct(req *Request, act *action.Action, schemata schema.Schemata) (Qualifications, error) {
	var ret Qualifications

	switch req.Type {
	case action.TypeModel.Key:
		srcKey, err := req.Params.GetString(action.TypeSource.Key, false)
		if err != nil {
			return nil, err
		}
		modelPath, err := req.Params.GetString(action.TypeModel.Key, false)
		if err != nil {
			return nil, err
		}
		ret, err = qualifyModel(req, act, srcKey, util.Pkg(util.StringSplitAndTrim(modelPath, "/")), schemata)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.Errorf("unhandled qualify type [%s]", req.Type)
	}

	kids, err := Qualify(req, act.Children, schemata)
	if err != nil {
		return nil, err
	}
	ret = append(ret, kids...)
	return ret, nil
}
