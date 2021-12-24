package workspace

import (
	"strings"

	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/export"
	"github.com/kyleu/admini/app/schema/model"
	"github.com/kyleu/admini/views/vexport"
	"github.com/pkg/errors"
)

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
		return nil, errors.Errorf("unhandled path [%s]", strings.Join(additional, "/"))
	}
}
