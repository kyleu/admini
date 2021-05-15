package controller

import (
	"github.com/kyleu/admini/app/model"
)

func modelView(req *workspaceRequest, m *model.Model, idStrings []string) (string, error) {
	return modelDetail(req, m, idStrings, "v")
}
