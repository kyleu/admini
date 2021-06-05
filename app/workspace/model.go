package workspace

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/action"

	"github.com/kyleu/admini/app/export"
	"github.com/kyleu/admini/views/vexport"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/views/vaction"
	"github.com/pkg/errors"
)

func processModel(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, additional []string) (*Result, error) {
	if len(additional) == 0 {
		return processModelList(req, act, srcKey, m)
	}
	switch additional[0] {
	case "new":
		return processModelNew(req, act, srcKey, m)
	case "export":
		return processModelExport(req, act, m, additional[1:])
	case "v":
		return processModelView(req, act, srcKey, m, additional[1:])
	case "x":
		return processModelEdit(req, act, srcKey, m, additional[1:])
	default:
		return nil, errors.Errorf("unhandled model parameters [%s]", strings.Join(additional, "/"))
	}
}

func processModelNew(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model) (*Result, error) {
	_, ld, _, err := loaderFor(req, srcKey)
	if err != nil {
		return ErrResult(req, act, err)
	}

	x, err := ld.Default(m)
	if err != nil {
		return ErrResult(req, act, errors.Wrapf(err, "can't load [%s] defaults", m.Key))
	}

	page := &vaction.ModelNew{Req: req, Act: act, Model: m, Defaults: x}
	return NewResult("", nil, req, act, x, page), nil
}

func processModelExport(req *cutil.WorkspaceRequest, act *action.Action, m *model.Model, additional []string) (*Result, error) {
	bc := append(act.Path(), req.Path...)
	switch len(additional) {
	case 0:
		page := &vexport.List{Req: req, Act: act, Model: m}
		return NewResult("", bc, req, act, nil, page), nil
	case 2:
		f := export.AllFormats.Get(additional[0], additional[1])
		res, err := export.Model(m, f, nil)
		if err != nil {
			return nil, errors.Wrap(err, "unable to export")
		}
		page := &vexport.Detail{Format: f, Results: res}
		return NewResult("", bc, req, act, nil, page), nil
	default:
		return nil, errors.Errorf("unhandled path [%v]", strings.Join(additional, "/"))
	}
}

func processModelList(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model) (*Result, error) {
	switch m.Type {
	case model.TypeStruct:
		_, ld, params, err := loaderFor(req, srcKey)
		if err != nil {
			return ErrResult(req, act, err)
		}

		rs, err := ld.List(m, params)
		if err != nil {
			return ErrResult(req, act, errors.Wrapf(err, "unable to list model [%s]", m.Key))
		}
		page := &vaction.ModelList{Req: req, Act: act, Model: m, ParamSet: params, Result: rs}
		return NewResult("", nil, req, act, rs, page), nil
	case model.TypeEnum:
		page := &vaction.ModelEnum{Req: req, Act: act, Model: m}
		return NewResult("", nil, req, act, m, page), nil
	default:
		return nil, errors.Errorf("unhandled model type [%s]", m.Type.String())
	}
}

func processModelView(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, idStrings []string) (*Result, error) {
	_, ld, params, err := loaderFor(req, srcKey)
	if err != nil {
		return ErrResult(req, act, err)
	}

	data, err := getModel(m, idStrings, ld)
	if err != nil {
		return ErrResult(req, act, err)
	}

	page := &vaction.ModelView{Req: req, Act: act, Model: m, ParamSet: params, Result: data}
	idx := len(req.Path) - len(idStrings) - 1
	if idx < 0 {
		idx = 0
	}
	bc := append(append(act.Path(), req.Path[:idx]...), idStrings...)
	return NewResult("", bc, req, act, data, page), nil
}

func processModelEdit(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, idStrings []string) (*Result, error) {
	_, ld, params, err := loaderFor(req, srcKey)
	if err != nil {
		return ErrResult(req, act, err)
	}

	data, err := getModel(m, idStrings, ld)
	if err != nil {
		return ErrResult(req, act, err)
	}

	page := &vaction.ModelEdit{Req: req, Act: act, Model: m, ParamSet: params, Result: data}
	idx := len(req.Path) - len(idStrings) - 1
	if idx < 0 {
		idx = 0
	}
	idBC := make([]string, 0, len(idStrings))
	u := req.RouteAct(act, 1+len(idStrings), append([]string{`v`}, idStrings...)...)
	for _, idString := range idStrings {
		idBC = append(idBC, fmt.Sprintf("%s||%s", idString, u))
	}

	bc := append(append(append(act.Path(), req.Path[:idx]...), idBC...), "edit")
	return NewResult("", bc, req, act, data, page), nil
}

func getModel(m *model.Model, idStrings []string, ld loader.Loader) ([]interface{}, error) {
	ids := make([]interface{}, 0, len(idStrings))
	for _, x := range idStrings {
		ids = append(ids, x)
	}

	rs, err := ld.Get(m, ids)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to retrieve model [%s] with key [%s]", m.Key, strings.Join(idStrings, "/"))
	}
	switch len(rs.Data) {
	case 0:
		return nil, errors.Errorf("no [%s] found with key [%s]", m.Name(), strings.Join(idStrings, "/"))
	case 1:
		return rs.Data[0], nil
	default:
		return nil, errors.Errorf("multiple [%s] found with key [%s]", m.Name(), strings.Join(idStrings, "/"))
	}
}

func loaderFor(req *cutil.WorkspaceRequest, srcKey string) (*source.Source, loader.Loader, util.ParamSet, error) {
	s, err := req.Sources.GetWithError(srcKey)
	if err != nil {
		return nil, nil, nil, err
	}

	l, err := req.AS.Loaders.Get(s.Type, s.Key, s.Config)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "no loader available")
	}
	p := cutil.ParamSetFromRequest(req.Ctx)
	return s, l, p, nil
}
