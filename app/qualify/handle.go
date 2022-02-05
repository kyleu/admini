package qualify

import (
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/lib/schema/model"
)

func Handle(rel *model.Relationship, act *action.Action, wr *cutil.WorkspaceRequest, m *model.Model, result []interface{}) (Results, error) {
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

	ret := make(Results, 0, len(quals))
	for _, q := range quals {
		mdl, _, err := action.GetModel(q.Action, wr.Schemata)
		if err != nil {
			return nil, err
		}

		// TODO: pull and fill `d`
		d := make([]interface{}, 0, len(mdl.Fields))
		for range mdl.Fields {
			d = append(d, "TODO")
		}

		x := &Result{Action: q.Action, Icon: q.Icon, Path: q.Path, Model: mdl, Data: d}
		ret = append(ret, x)
	}

	return ret, nil
}
