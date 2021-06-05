package cutil

import (
	"fmt"
	"github.com/kyleu/admini/app/action"
	"path/filepath"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/source"
	"github.com/valyala/fasthttp"
)

type WorkspaceRequest struct {
	T        string               `json:"t"`
	K        string               `json:"k"`
	Ctx      *fasthttp.RequestCtx `json:"-"`
	AS       *app.State           `json:"-"`
	PS       *PageState           `json:"-"`
	Item     interface{}          `json:"item,omitempty"`
	Path     []string             `json:"path,omitempty"`
	Project  *project.Project     `json:"-"`
	Sources  source.Sources       `json:"-"`
	Schemata schema.Schemata      `json:"-"`
}

func (r *WorkspaceRequest) Route(path ...string) string {
	url := fmt.Sprintf("/%s/%s", r.T, r.K)
	if len(path) > 0 {
		url += "/" + filepath.Join(path...)
	}
	return url
}

func (r *WorkspaceRequest) RouteAct(act *action.Action, drop int, path ...string) string {
	end := len(r.Path) - drop
	if end < 0 {
		end = 0
	}
	return r.Route(append(append(act.Path(), r.Path[:end]...), path...)...)
}
