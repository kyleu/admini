package controller

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/views/vutil"
)

func modelEdit(req *workspaceRequest, m *model.Model, idStrings []string) (string, error) {
	return modelLink(req, m, idStrings, "x")
}

func modelSave(req *workspaceRequest, m *model.Model, idStrings []string) (string, error) {
	println(strings.Join(req.Path, "/"))
	changes, err := cutil.ParseForm(req.R)
	if err != nil {
		return ersp("unable to parse changes: %w", err)
	}

	msg := fmt.Sprintf("saved [%v] changes to %v [%v]", len(changes), m.Key, strings.Join(idStrings, "/"))
	url := vutil.WorkspaceLink(req.AS, req.T, req.K, append(append(m.Path(), "v"), idStrings...)...)
	return flashAndRedir(true, msg, url, req.W, req.R, req.PS)
}
