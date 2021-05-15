package controller

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/views/layout"

	"github.com/kyleu/admini/views/vworkspace"
)

func handleModel(req *workspaceRequest, m *model.Model) (string, error) {
	if len(req.Path) == 0 {
		return modelList(req, m)
	}
	switch req.Path[0] {
	case "v":
		return modelView(req, m, req.Path[1:])
	case "x":
		switch req.R.Method {
		case http.MethodGet:
			return modelEdit(req, m, req.Path[1:])
		case http.MethodPost:
			return modelSave(req, m, req.Path[1:])
		default:
			return whoops(req, "unhandled method ["+req.R.Method+"]", m.Path()...)
		}
	case "export":
		return modelExport(req, m)
	default:
		return whoops(req, "unhandled model action", append(m.Path(), req.Path...)...)
	}
}

func modelList(req *workspaceRequest, m *model.Model) (string, error) {
	params := cutil.ParamSetFromRequest(req.R)
	l := req.AS.Loaders.Get(req.Src.Type)
	if l == nil {
		return ersp("no loader [" + req.Src.Type.String() + "] available")
	}

	rs, err := l.List(req.Src.Key, req.Src.Config, m, params)
	if err != nil {
		return "", errors.Wrap(err, "unable to list model ["+m.Key+"]")
	}

	req.PS.Data = rs

	page := &vworkspace.ModelList{Model: m, ParamSet: params, CtxT: req.T, CtxK: req.K, Result: rs}
	return render(req.R, req.W, req.AS, page, req.PS, m.Path()...)
}

func modelDetail(req *workspaceRequest, m *model.Model, idStrings []string, act string) (string, error) {
	l := req.AS.Loaders.Get(req.Src.Type)
	if l == nil {
		return ersp("no loader [" + req.Src.Type.String() + "] available")
	}

	ids := make([]interface{}, 0, len(idStrings)-1)
	for _, x := range req.Path[1:] {
		ids = append(ids, x)
	}

	rs, err := l.Get(req.Src.Key, req.Src.Config, m, ids)
	if err != nil {
		return "", errors.Wrap(err, "unable to retrieve model ["+m.Key+"]")
	}

	req.PS.Data = rs
	var page layout.Page

	bc := append(m.Path(), idStrings...)

	switch rs.Size() {
	case 0:
		return whoops(req, "no model found with id ["+strings.Join(idStrings, "/")+"]", append(m.Path(), req.Path[1:]...)...)
	case 1:
		switch act {
		case "v":
			page = &vworkspace.ModelView{Model: m, CtxT: req.T, CtxK: req.K, Result: rs.Data[0]}
			return render(req.R, req.W, req.AS, page, req.PS, bc...)
		case "x":
			bc = append(bc, "edit")
			page = &vworkspace.ModelEdit{Model: m, CtxT: req.T, CtxK: req.K, Result: rs.Data[0]}
			return render(req.R, req.W, req.AS, page, req.PS, bc...)
		default:
			return whoops(req, "unhandled action ["+act+"]", append(m.Path(), idStrings...)...)
		}
	default:
		return whoops(req, "multiple models found with id ["+strings.Join(idStrings, "/")+"]")
	}
}
