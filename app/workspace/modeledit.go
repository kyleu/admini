package workspace

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/views/vmodel"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
	"github.com/pkg/errors"
)

func processModelEdit(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, idStrings []string) (*Result, error) {
	_, ld, params, err := loaderFor(req, srcKey)
	if err != nil {
		return ErrResult(req, act, err)
	}

	data, err := getModel(m, idStrings, ld)
	if err != nil {
		return ErrResult(req, act, err)
	}

	page := &vmodel.Edit{Req: req, Act: act, Model: m, ParamSet: params, Result: data}
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

func processModelSave(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, idStrings []string) (*Result, error) {
	changes, err := cutil.ParseFormAsChanges(req.Ctx)
	if err != nil {
		return ErrResult(req, act, err)
	}

	_, ld, _, err := loaderFor(req, srcKey)
	if err != nil {
		return ErrResult(req, act, err)
	}

	res, err := ld.Save(m, changes)
	if err != nil {
		return ErrResult(req, act, err)
	}

	return nil, errors.Errorf("Saved [%d] changes, received [%d] fields", len(changes), len(res))
}

func processModelDelete(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, idStrings []string) (*Result, error) {
	pk := m.GetPK(req.PS.Logger)

	_, ld, _, err := loaderFor(req, srcKey)
	if err != nil {
		return ErrResult(req, act, err)
	}

	vals := make([]interface{}, 0, len(idStrings))
	for _, s := range idStrings {
		vals = append(vals, s)
	}

	numAffected, err := ld.Remove(m, pk, vals, 1)
	if err != nil {
		return ErrResult(req, act, err)
	}
	if numAffected != 1 {
		return ErrResult(req, act, errors.Errorf("expected one deleted row, observed [%d]", numAffected))
	}

	msg := fmt.Sprintf("Deleted %s [%s]", m.Name(), strings.Join(idStrings, ":"))
	return RedirectResult(msg, req.Route(act.Path()...)), nil
}
