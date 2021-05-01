package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kyleu/admini/app/source"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/views"
	"github.com/kyleu/admini/views/vworkspace"

	"github.com/kyleu/admini/app"
)

func handleModel(w http.ResponseWriter, r *http.Request, as *app.State, ps *cutil.PageState, src *source.Source, m *schema.Model, remaining []string) (string, error) {
	ps.Data = m
	if len(remaining) == 0 {
		return modelList(w, r, as, ps, src, m)
	}
	page := &views.TODO{Message: fmt.Sprintf("unhandled model action [%v]", strings.Join(remaining, "/"))}
	return render(r, w, as, page, ps, m.Path()...)
}

func modelList(w http.ResponseWriter, r *http.Request, as *app.State, ps *cutil.PageState, src *source.Source, m *schema.Model) (string, error) {
	params := cutil.ParamSetFromRequest(r)
	l := as.Loaders.Get(src.Type)
	if l == nil {
		return ersp("no loader [" + src.Type.String() + "] available")
	}

	rs, err := l.List(src.Key, src.Config, m, params)
	if err != nil {
		return ersp("unable to list model ["+m.Key+"]: %w", err)
	}

	page := &vworkspace.ModelList{Model: m, ParamSet: params, Source: src.Key, Result: rs}
	return render(r, w, as, page, ps, m.Path()...)
}
