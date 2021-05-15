package controller

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/schema"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/views/vutil"
)

func modelEdit(req *workspaceRequest, m *model.Model, idStrings []string) (string, error) {
	return modelDetail(req, m, idStrings, "x")
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

	msg := fmt.Sprintf("saved [%v] changes to %v [%v]", len(changes), m.Key, strings.Join(idStrings, "/"))
	url := vutil.WorkspaceLink(req.AS, req.T, req.K, append(append(m.Path(), "v"), idStrings...)...)

	ids := make([]interface{}, 0, len(idStrings)-1)
	for _, x := range req.Path[1:] {
		ids = append(ids, x)
	}

	ld := req.AS.Loaders.Get(schema.OriginPostgres)

	curr, err := ld.Get(req.Src.Key, req.Src.Config, m, ids)
	if err != nil {
		return "", errors.Wrap(err, "unable to parse changes")
	}

	if curr.Size() == 0 {
		return ersp("can't load original [%v] with id [%v]", m.Key, strings.Join(idStrings, " / "))
	}
	if curr.Size() > 1 {
		return ersp("multiple [%v] matched key [%v]", m.Key, strings.Join(idStrings, " / "))
	}
	data := curr.Data[0]

	for k, v := range changes {
		idx, _ := curr.Fields.Get(k)
		orig := data[idx]
		if orig == v {
			println(fmt.Sprintf("MATCH [%v]: %v == %v", k, orig, v))
		} else {
			println(fmt.Sprintf("NO MATCH [%v]: %v != %v", k, orig, v))
		}
	}

	return flashAndRedir(true, msg, url, req.W, req.R, req.PS)
}
