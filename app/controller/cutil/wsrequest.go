package cutil

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"

	"admini.dev/admini/app/action"
	"admini.dev/admini/app/lib/schema"
	"admini.dev/admini/app/project"
	"admini.dev/admini/app/source"
)

type WorkspaceRequest struct {
	T        string              `json:"t"`
	K        string              `json:"k"`
	Req      *http.Request       `json:"-"`
	ReqBody  []byte              `json:"-"`
	Rsp      http.ResponseWriter `json:"-"`
	PS       *PageState          `json:"-"`
	Item     any                 `json:"item,omitempty"`
	Path     []string            `json:"path,omitempty"`
	Project  *project.Project    `json:"-"`
	Sources  source.Sources      `json:"-"`
	Schemata schema.Schemata     `json:"-"`
	Context  context.Context     `json:"-"` //nolint:containedctx // properly closed, never directly used
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

func (r *WorkspaceRequest) Clone() *WorkspaceRequest {
	return &WorkspaceRequest{
		T: r.T, K: r.K, Req: r.Req, ReqBody: r.ReqBody, Rsp: r.Rsp, PS: r.PS,
		Item: r.Item, Path: r.Path, Project: r.Project,
		Sources: r.Sources, Schemata: r.Schemata, Context: r.Context,
	}
}
