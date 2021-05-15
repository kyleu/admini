package controller

import (
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/sandbox/export"
	"github.com/kyleu/admini/views/vsandbox"
)

func modelExport(req *workspaceRequest, m *model.Model) (string, error) {
	out, err := export.Model(m, req.PS.Logger)
	if err != nil {
		return whoops(req, "unable to export", append(m.Path(), req.Path...)...)
	}
	req.PS.Data = out
	page := &vsandbox.Export{Title: m.Key, Results: out}
	return render(req.R, req.W, req.AS, page, req.PS, m.Path()...)
}
