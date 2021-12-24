package qualify

import (
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	model2 "github.com/kyleu/admini/app/schema/model"
)

func Handle(rel *model2.Relationship, act *action.Action, wr *cutil.WorkspaceRequest, m *model2.Model, result []interface{}) (Results, error) {
	rowFK, err := model2.GetStrings(m.Fields, rel.SourceFields, result)
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
		// TODO: pull and fill `m` && `d`

		mdl := &model2.Model{Key: "TODO"}
		d := []interface{}{"TODO"}
		x := &Result{Action: q.Action, Icon: q.Icon, Path: q.Path, Model: mdl, Data: d}
		ret = append(ret, x)
	}

	return ret, nil
}
