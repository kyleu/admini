package workspace

import (
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/views/vaction"
)

func processPackage(req *cutil.WorkspaceRequest, act *action.Action, t *model.Package) (*Result, error) {
	page := &vaction.Package{Req: req, Act: act, Pkg: t}
	return NewResult("", nil, req, act, t, page), nil
}
