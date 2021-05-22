package controller

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/model"

	"github.com/kyleu/admini/views"
	"github.com/kyleu/admini/views/vworkspace"
)

func handlePackage(req *workspaceRequest, mp *model.Package) (string, error) {
	req.PS.Title = mp.PathString()
	req.PS.Data = mp
	if len(req.Path) == 0 {
		page := &vworkspace.PackageDetail{Pkg: mp, CtxT: req.T, CtxK: req.K}
		return render(req.R, req.W, req.AS, page, req.PS, mp.Path()...)
	}
	page := &views.TODO{Message: fmt.Sprintf("unhandled package action [%v]", strings.Join(req.Path, "/"))}
	return render(req.R, req.W, req.AS, page, req.PS, mp.Path()...)
}
