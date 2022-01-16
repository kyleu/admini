package workspace

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/lib/schema/model"
	"github.com/kyleu/admini/views/vmodel"
)

func processModelEdit(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, idStrings []string, as *app.State) (*Result, error) {
	_, ld, err := loaderFor(req, srcKey, as)
	if err != nil {
		return ErrResult(req, act, err)
	}

	data, err := getModel(req.Context, m, idStrings, ld)
	if err != nil {
		return ErrResult(req, act, err)
	}

	page := &vmodel.Edit{Req: req, Act: act, Model: m, Result: data}
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

func processModelSave(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, idStrings []string, as *app.State) (*Result, error) {
	changes, err := cutil.ParseFormAsChanges(req.Ctx)
	if err != nil {
		return ErrResult(req, act, err)
	}

	if len(changes) == 0 {
		return RedirectResult("no changes required", req.RouteAct(act, 0)), nil
	}

	_, ld, err := loaderFor(req, srcKey, as)
	if err != nil {
		return ErrResult(req, act, err)
	}

	ids := make([]interface{}, 0, len(idStrings))
	for _, s := range idStrings {
		ids = append(ids, s)
	}

	res, err := ld.Save(req.Context, m, ids, changes)
	if err != nil {
		return ErrResult(req, act, err)
	}

	msg := fmt.Sprintf("Added [%d] changes, received [%d] fields", len(changes), len(res))
	path := []string{"v"}
	for _, x := range res {
		path = append(path, fmt.Sprintf("%v", x))
	}
	dest := req.RouteAct(act, len(m.GetPK(req.PS.Logger))+1, path...)
	return RedirectResult(msg, dest), nil
}

func processModelDelete(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, idStrings []string, as *app.State) (*Result, error) {
	pk := m.GetPK(req.PS.Logger)

	_, ld, err := loaderFor(req, srcKey, as)
	if err != nil {
		return ErrResult(req, act, err)
	}

	vals := make([]interface{}, 0, len(idStrings))
	for _, s := range idStrings {
		vals = append(vals, s)
	}

	numAffected, err := ld.Remove(req.Context, m, pk, vals, 1)
	if err != nil {
		return ErrResult(req, act, err)
	}
	if numAffected != 1 {
		return ErrResult(req, act, errors.Errorf("expected one deleted row, observed [%d]", numAffected))
	}

	msg := fmt.Sprintf("Deleted %s [%s]", m.Name(), strings.Join(idStrings, ":"))
	return RedirectResult(msg, req.RouteAct(act, len(idStrings)+1)), nil
}
