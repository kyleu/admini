package controller

import (
	"fmt"
	"net/http"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vproject"

	"github.com/gorilla/mux"
)

func ProjectList(w http.ResponseWriter, r *http.Request) {
	act("project.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		p, err := as.Projects.List()
		if err != nil {
			return "", fmt.Errorf("unable to load project list: %w", err)
		}
		ps.Data = p
		return render(r, w, as, &vproject.ProjectList{Projects: p}, ps, "projects")
	})
}

func ProjectDetail(w http.ResponseWriter, r *http.Request) {
	act("project.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		prj, err := as.Projects.Load(key)
		if err != nil {
			return "", fmt.Errorf("unable to load project ["+key+"]: %w", err)
		}
		ps.Data = prj
		return render(r, w, as, &vproject.ProjectDetail{Project: prj}, ps, "projects", prj.Key)
	})
}
