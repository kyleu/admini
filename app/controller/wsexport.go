package controller

import (
	"fmt"
	"github.com/kyleu/admini/app/export"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/views/vexport"
)

func modelExport(req *workspaceRequest, m *model.Model, keys []string) (string, error) {
	req.PS.Title = fmt.Sprintf("Export [%v]", m.Name())
	if len(keys) == 0 {
		req.PS.Title = "Export"
		req.PS.Data = export.AllFormats
		page := &vexport.List{CtxT: req.T, CtxK: req.K, Model: m}
		return render(req.R, req.W, req.AS, page, req.PS, append(m.Path(), "export")...)
	}
	if len(keys) != 2 {
		return ersp("must provide language and flavor")
	}
	f := export.AllFormats.Get(keys[0], keys[1])
	if f == nil {
		return ersp("invalid export type [" + keys[0] + "/" + keys[1] + "]")
	}
	out, err := export.Model(m, f, req.PS.Logger)
	if err != nil {
		return whoops(req, "unable to export", append(m.Path(), req.Path...)...)
	}
	req.PS.Title = "Export Result"
	req.PS.Data = out
	page := &vexport.Detail{Format: f, Results: out}
	return render(req.R, req.W, req.AS, page, req.PS, append(m.Path(), "export")...)
}
