package workspace

import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/export"
	"github.com/kyleu/admini/app/schema/model"
	"github.com/kyleu/admini/views/vexport"
	"github.com/pkg/errors"
)

func processModelExport(req *cutil.WorkspaceRequest, act *action.Action, m *model.Model, additional []string, as *app.State) (*Result, error) {
	bc := append(act.Path(), req.Path...)
	res, err := export.Export(m, as.Logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to export")
	}
	page := &vexport.Detail{Results: res}
	return NewResult("", bc, req, act, nil, page), nil
}
