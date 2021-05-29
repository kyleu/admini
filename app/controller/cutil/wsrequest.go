package cutil

import (
	"net/http"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/project/action"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/views/vutil"
)

type WorkspaceRequest struct {
	T        string              `json:"t"`
	K        string              `json:"k"`
	W        http.ResponseWriter `json:"-"`
	R        *http.Request       `json:"-"`
	AS       *app.State          `json:"-"`
	PS       *PageState          `json:"-"`
	Item     interface{}         `json:"item,omitempty"`
	Path     []string            `json:"path,omitempty"`
	Project  *project.Project    `json:"-"`
	Sources  source.Sources      `json:"-"`
	Schemata schema.Schemata     `json:"-"`
}

func (r *WorkspaceRequest) Route(path ...string) string {
	url := vutil.WorkspaceLink(r.AS, r.T, r.K, path...)
	return url
}

func (r *WorkspaceRequest) RouteAct(act *action.Action, drop int, path ...string) string {
	end := len(r.Path) - drop
	if end < 0 {
		end = 0
	}
	return r.Route(append(append(act.Path(), r.Path[:end]...), path...)...)
}
