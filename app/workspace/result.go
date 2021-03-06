package workspace

import (
	"fmt"

	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/views/layout"
)

type Result struct {
	Title       string      `json:"title"`
	Redirect    string      `json:"redirect,omitempty"`
	Breadcrumbs []string    `json:"bc"`
	Data        any         `json:"data"`
	Page        layout.Page `json:"page"`
}

func NewResult(title string, bc []string, req *cutil.WorkspaceRequest, act *action.Action, data any, page layout.Page) *Result {
	if title == "" && act != nil {
		title = act.Name()
	}
	if bc == nil && act != nil {
		bc = append(act.Path(), req.Path...)
	}
	return &Result{Title: title, Breadcrumbs: bc, Data: data, Page: page}
}

func RedirectResult(flashMessage string, redir string) *Result {
	return &Result{Title: flashMessage, Redirect: redir}
}

func ErrResult(req *cutil.WorkspaceRequest, act *action.Action, err error) (*Result, error) {
	title := "Error"
	if act != nil {
		title = fmt.Sprintf("%s: %s", title, act.Name())
	}
	return NewResult(title, nil, req, act, fmt.Sprintf("ERROR: %s", err.Error()), nil), err
}
