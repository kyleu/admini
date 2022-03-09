package workspace

import (
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/views/vaction"
)

func processPackage(req *cutil.WorkspaceRequest, act *action.Action, t *model.Package) (*Result, error) {
	page := &vaction.Package{Req: req, Act: act, Pkg: t}
	return NewResult("", nil, req, act, t, page), nil
}
