package controller

import (
	"net/http"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/project/action"
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
		prj, err := as.Projects.LoadRequired(key, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to load project ["+key+"]")
		}
		ps.Title = prj.Name()
		ps.Data = prj
		return render(r, w, as, &vproject.Detail{Project: prj}, ps, "projects", prj.Key)
	})
}

func ProjectTest(w http.ResponseWriter, r *http.Request) {
	act("project.test", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		prj, err := as.Projects.LoadRequired(key, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to load project ["+key+"]")
		}

		_, err = action.Save(prj.Key, prj.Actions, as.Files)
		if err != nil {
			return "", errors.Wrap(err, "unable to save actions for project ["+key+"]")
		}

		_, err = currentApp.Projects.LoadRequired(key, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to reload project ["+key+"]")
		}

		ps.Title = "Test " + prj.Name()
		ps.Data = prj
		return render(r, w, as, &views.TODO{Message: "OK!"}, ps, "projects", prj.Key, "test")
	})
}
