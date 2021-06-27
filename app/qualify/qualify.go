package qualify

import (
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
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

func Handle(rel *model.Relationship, act *action.Action, wr *cutil.WorkspaceRequest, m *model.Model, result []interface{}) (Qualifications, error) {
	rowFK, err := model.GetStrings(m.Fields, rel.SourceFields, result)
	if err != nil {
		return nil, err
	}
	src := act.Config["source"]
	if act.Type == action.TypeAll {
		src = wr.Path[0]
	}
	req := NewRequest("model", "view", "source", src, "model", rel.Path(), "keys", rowFK)

	return Qualify(req, wr.Project.Actions, wr.Schemata)
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
		ret, err = qualifyModel(req, act, srcKey, util.SplitAndTrim(modelPath, "/"), schemata)
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
