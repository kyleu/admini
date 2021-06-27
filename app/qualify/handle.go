package qualify

import (
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
)

func Handle(rel *model.Relationship, act *action.Action, wr *cutil.WorkspaceRequest, m *model.Model, result []interface{}) (Results, error) {
	rowFK, err := model.GetStrings(m.Fields, rel.SourceFields, result)
	if err != nil {
		return nil, err
	}
	src := act.Config["source"]
	if act.Type == action.TypeAll {
		src = wr.Path[0]
	}
	req := NewRequest("model", "view", "source", src, "model", rel.Path(), "keys", rowFK)

	quals, err := Qualify(req, wr.Project.Actions, wr.Schemata)
	if err != nil {
		return nil, err
	}

	ret := make(Results, 0, len(quals))
	for _, q := range quals {
		// pull and fill `m` && `d`
		var m *model.Model
		var d []interface{}
		x := &Result{Action: q.Action, Icon: q.Icon, Path: q.Path, Model: m, Data: d}
		ret = append(ret, x)
	}

	return ret, nil
}
