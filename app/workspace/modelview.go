package workspace

import (
	"strings"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/lib/schema/model"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/views/vmodel"
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
		//TODO maybe: getModel(req.Context, someModel, idStrings, ld)
		req.PS.Logger.Debugf("relations: [%s]", strings.Join(rowFK, ", "))
	}

	obj := util.NewOrderedMap(false, len(m.Fields))
	for idx, f := range m.Fields {
		obj.Append(f.Key, data[idx])
	}

	page := &vmodel.View{Req: req, Act: act, Model: m, Result: data}
	idx := len(req.Path) - len(idStrings) - 1
	if idx < 0 {
		idx = 0
	}
	bc := append(append(act.Path(), req.Path[:idx]...), idStrings...)
	return NewResult("", bc, req, act, obj, page), nil
}
