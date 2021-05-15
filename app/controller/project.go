package controller

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vproject"

	"github.com/gorilla/mux"
)

func ProjectList(w http.ResponseWriter, r *http.Request) {
	act("project.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		p, err := as.Projects.List()
		if err != nil {
			return "", errors.Wrap(err, "unable to load project list")
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
			return "", errors.Wrap(err, "unable to load project ["+key+"]")
		}
		ps.Data = prj
		return render(r, w, as, &vproject.ProjectDetail{Project: prj}, ps, "projects", prj.Key)
	})
}
