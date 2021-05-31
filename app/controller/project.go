package controller

import (
	"net/http"

	"github.com/kyleu/admini/app/project/action"
	"github.com/kyleu/admini/app/util"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/views"
	"github.com/pkg/errors"

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
		ps.Title = "Projects"
		ps.Data = p
		return render(r, w, as, &vproject.List{Projects: p}, ps, "projects")
	})
}

func ProjectDetail(w http.ResponseWriter, r *http.Request) {
	act("project.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		prj, err := as.Projects.LoadView(key)
		if err != nil {
			return "", errors.Wrap(err, "unable to load project ["+key+"]")
		}
		ps.Title = prj.Project.Name()
		ps.Data = prj.Project
		return render(r, w, as, &vproject.Detail{View: prj}, ps, "projects", prj.Project.Key)
	})
}

func ProjectTest(w http.ResponseWriter, r *http.Request) {
	act("project.test", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		v, err := as.Projects.LoadView(key)
		if err != nil {
			return "", errors.Wrap(err, "unable to load project ["+key+"]")
		}

		req := action.NewRequest(action.TypeModel.Key, "list", action.TypeSource.Key, "admini_test", action.TypeModel.Key, "public/simple")

		q, err := action.Qualify(req, v.Project.Actions)
		if err != nil {
			return "", errors.Wrap(err, "unable to qualify project ["+key+"]")
		}

		ps.Title = v.Project.Name() + " - Test"
		ps.Data = util.ValueMap{
			//"project": v.Project,
			"request": req,
			"qualify": q,
		}

		view := &views.TODO{Message: "Project [" + v.Project.Key + "]: OK"}
		return render(r, w, as, view, ps, "projects", v.Project.Key, "test")
	})
}
