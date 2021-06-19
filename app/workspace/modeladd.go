package workspace

import (
	"fmt"

	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/views/vmodel"
	"github.com/pkg/errors"
)

func processModelNew(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model) (*Result, error) {
	_, ld, _, err := loaderFor(req, srcKey)
	if err != nil {
		return ErrResult(req, act, err)
	}

	x, err := ld.Default(m)
	if err != nil {
		return ErrResult(req, act, errors.Wrapf(err, "can't load [%s] defaults", m.Key))
	}

	page := &vmodel.New{Req: req, Act: act, Model: m, Defaults: x}
	return NewResult("", nil, req, act, x, page), nil
}

func processModelAdd(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, additional []string) (*Result, error) {
	changes, err := cutil.ParseFormAsChanges(req.Ctx)
	if err != nil {
		return ErrResult(req, act, err)
	}

	_, ld, _, err := loaderFor(req, srcKey)
	if err != nil {
		return ErrResult(req, act, err)
	}

	res, err := ld.Add(m, changes)
	if err != nil {
		return ErrResult(req, act, err)
	}

	msg := fmt.Sprintf("Added [%d] changes, received [%d] fields", len(changes), len(res))
	return RedirectResult(msg, req.RouteAct(act, 1, additional[:len(additional)-1]...)), nil
}
