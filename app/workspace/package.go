package workspace

import (
	"admini.dev/app/action"
	"admini.dev/app/controller/cutil"
	"admini.dev/app/lib/schema/model"
	"admini.dev/views/vaction"
)

func processPackage(req *cutil.WorkspaceRequest, act *action.Action, t *model.Package) (*Result, error) {
	page := &vaction.Package{Req: req, Act: act, Pkg: t}
	return NewResult("", nil, req, act, t, page), nil
}
