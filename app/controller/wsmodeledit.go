package controller

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/views/vworkspace"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/views/vutil"
)

func modelEdit(req *workspaceRequest, m *model.Model, idStrings []string) (string, error) {
	return modelDetail(req, m, idStrings, "x")
}

func modelNew(req *workspaceRequest, m *model.Model) (string, error) {
	l, err := req.AS.Loaders.Get(req.Src.Type, req.Src.Key, req.Src.Config)
	if err != nil {
		return "", errors.Wrap(err, "no loader available")
	}

	x, err := l.Default(m)
	if err != nil {
		return "", errors.Wrap(err, "can't load ["+m.Key+"] defaults")
	}

	page := &vworkspace.ModelNew{CtxT: req.T, CtxK: req.K, Model: m, Defaults: x}
	return render(req.R, req.W, req.AS, page, req.PS, append(m.Path(), "new")...)
}

func modelSave(req *workspaceRequest, m *model.Model, idStrings []string) (string, error) {
	form, err := cutil.ParseForm(req.R)
	if err != nil {
		return "", errors.Wrap(err, "unable to parse form")
	}
	changes, err := form.AsChanges()
	if err != nil {
		return "", errors.Wrap(err, "unable to parse changes")
	}

	ld, err := req.AS.Loaders.Get(req.Src.Type, req.Src.Key, req.Src.Config)
	if err != nil {
		return "", errors.Wrap(err, "unable to create loader")
	}

	if len(idStrings) == 0 {
		return wsinsert(req, m, changes, ld)
	}
	return wsupdate(req, m, changes, ld, idStrings)
}

func wsinsert(req *workspaceRequest, m *model.Model, changes map[string]interface{}, ld loader.Loader) (string, error) {
	pkDef := m.GetPK(req.PS.Logger)
	idStrings := make([]string, 0, len(pkDef))
	for _, x := range pkDef {
		idStrings = append(idStrings, fmt.Sprintf("%v", changes[x]))
	}

	msg := fmt.Sprintf("added new %v [%v] with [%v] fields", m.Name(), strings.Join(idStrings, "/"), len(changes))
	url := vutil.WorkspaceLink(req.AS, req.T, req.K, append(append(m.Path(), "v"), idStrings...)...)

	return flashAndRedir(true, msg, url, req.W, req.R, req.PS)
}

func wsupdate(req *workspaceRequest, m *model.Model, changes map[string]interface{}, ld loader.Loader, idStrings []string) (string, error) {
	ids := make([]interface{}, 0, len(idStrings))
	for _, x := range req.Path[1:] {
		ids = append(ids, x)
	}

	curr, err := ld.Get(m, ids)
	if err != nil {
		return "", errors.Wrap(err, "unable to parse changes")
	}

	if curr.Size() == 0 {
		return ersp("can't load original [%v] with id [%v]", m.Key, strings.Join(idStrings, " / "))
	}
	if curr.Size() > 1 {
		return ersp("multiple [%v] matched key [%v]", m.Key, strings.Join(idStrings, " / "))
	}
	fields, data, err := curr.SingleRow()
	if err != nil {
		return "", errors.Wrap(err, "return didn't have single row")
	}

	for k, v := range changes {
		idx, _ := fields.Get(k)
		orig := data[idx]
		if orig == v {
			println(fmt.Sprintf("MATCH [%v]: %v == %v", k, orig, v))
		} else {
			println(fmt.Sprintf("NO MATCH [%v]: %v != %v", k, orig, v))
		}
	}

	msg := fmt.Sprintf("saved [%v] changes to %v [%v]", len(changes), m.Name(), strings.Join(idStrings, "/"))
	url := vutil.WorkspaceLink(req.AS, req.T, req.K, append(append(m.Path(), "v"), idStrings...)...)

	return flashAndRedir(true, msg, url, req.W, req.R, req.PS)
}
