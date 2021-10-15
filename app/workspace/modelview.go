package workspace

import (
	"strings"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/views/vmodel"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
)

func processModelView(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, idStrings []string, as *app.State) (*Result, error) {
	_, ld, err := loaderFor(req, srcKey, as)
	if err != nil {
		return ErrResult(req, act, err)
	}

	data, err := getModel(req.Context, m, idStrings, ld)
	if err != nil {
		return ErrResult(req, act, err)
	}

	for _, rel := range m.Relationships {
		rowFK, err := model.GetStrings(m.Fields, rel.SourceFields, data)
		if err != nil {
			return nil, err
		}
		println("[" + strings.Join(rowFK, ", ") + "]")
	}

	page := &vmodel.View{Req: req, Act: act, Model: m, Result: data}
	idx := len(req.Path) - len(idStrings) - 1
	if idx < 0 {
		idx = 0
	}
	bc := append(append(act.Path(), req.Path[:idx]...), idStrings...)
	return NewResult("", bc, req, act, data, page), nil
}
