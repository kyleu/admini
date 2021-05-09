package controller

import (
	"fmt"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
	"strings"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/views/vutil"
)

func modelEdit(req *workspaceRequest, m *model.Model, idStrings []string) (string, error) {
	return modelLink(req, m, idStrings, "x")
}

const sfx = "--selected"

func modelSave(req *workspaceRequest, m *model.Model, idStrings []string) (string, error) {
	form, err := cutil.ParseForm(req.R)
	if err != nil {
		return "", errors.Wrap(err, "unable to parse form")
	}
	changes, err := parseChanges(form)
	if err != nil {
		return "", errors.Wrap(err, "unable to parse changes")
	}

	msg := fmt.Sprintf("saved [%v] changes to %v [%v]", len(changes), m.Key, strings.Join(idStrings, "/"))
	url := vutil.WorkspaceLink(req.AS, req.T, req.K, append(append(m.Path(), "v"), idStrings...)...)

	println(fmt.Sprintf("EDIT:\n%v\n%v\n%v", msg, url, util.ToJSON(changes)))

	return flashAndRedir(true, msg, url, req.W, req.R, req.PS)
}

func parseChanges(changes cutil.FormValues) (map[string]interface{}, error) {
	keys := []string{}
	vals := map[string]interface{}{}

	for _, f := range changes {
		if strings.HasSuffix(f.Key, sfx) {
			k := strings.TrimSuffix(f.Key, sfx)
			keys = append(keys, k)
		} else {
			curr, ok := vals[f.Key]
			if ok {
				return nil, errors.New(fmt.Sprintf("multiple values presented for [%v] (%v/%v)", f.Key, curr, f.Value))
			}
			vals[f.Key] = f.Value
		}
	}

	ret := make(map[string]interface{}, len(keys))
	for _, k := range keys {
		ret[k] = vals[k]
	}
	return ret, nil
}
