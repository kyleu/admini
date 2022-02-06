package workspace

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/lib/schema/model"
	"github.com/kyleu/admini/views/vmodel"
)

func processModelNew(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, as *app.State) (*Result, error) {
	_, ld, err := loaderFor(req, srcKey, as)
	if err != nil {
		return ErrResult(req, act, err)
	}

	x, err := ld.Default(req.Context, m)
	if err != nil {
		return ErrResult(req, act, errors.Wrapf(err, "can't load [%s] defaults", m.Key))
	}

	page := &vmodel.New{Req: req, Act: act, Model: m, Defaults: x}
	return NewResult("", nil, req, act, x, page), nil
}

func processModelAdd(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, additional []string, as *app.State) (*Result, error) {
	changes, err := cutil.ParseFormAsChanges(req.RC)
	if err != nil {
		return ErrResult(req, act, err)
	}

	_, ld, err := loaderFor(req, srcKey, as)
	if err != nil {
		return ErrResult(req, act, err)
	}

	res, err := ld.Add(req.Context, m, changes)
	if err != nil {
		return ErrResult(req, act, err)
	}

	msg := fmt.Sprintf("Added [%d] changes", len(changes))
	end := len(additional)
	if end > 0 {
		end--
	}
	path := make([]string, 0, end+len(res))
	path = append(path, additional[:end]...)
	path = append(path, "v")
	for _, x := range res {
		path = append(path, fmt.Sprintf("%v", x))
	}
	return RedirectResult(msg, req.RouteAct(act, 1, path...)), nil
}
