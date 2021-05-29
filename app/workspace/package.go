package workspace

import (
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/project/action"
	"github.com/kyleu/admini/views/vaction"
)

func processPackage(req *cutil.WorkspaceRequest, act *action.Action, root *model.Package, path []string, t *model.Package, remaining []string) (*Result, error) {
	page := &vaction.ResultPackage{Req: req, Act: act, Pkg: t}
	return NewResult("", nil, req, act, t, page), nil
}
