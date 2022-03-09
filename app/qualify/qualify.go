package qualify

import (
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/schema/model"
	"github.com/pkg/errors"

	"admini.dev/admini/app/action"
	"admini.dev/admini/app/lib/schema"
	"admini.dev/admini/app/util"
)

func Handle(rel *model.Relationship, act *action.Action, wr *cutil.WorkspaceRequest, m *model.Model, result []interface{}) (Qualifications, error) {
	rowFK, err := model.GetStrings(m.Fields, rel.SourceFields, result)
	if err != nil {
		return nil, err
	}
	src := act.Config["source"]
	if act.TypeKey == action.TypeAll.Key {
		src = wr.Path[0]
	}
	req := NewRequest("model", "view", "source", src, "model", rel.Path(), "keys", rowFK)

	quals, err := Qualify(req, wr.Project.Actions, wr.Schemata)
	if err != nil {
		return nil, err
	}

	return quals, nil
}

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
