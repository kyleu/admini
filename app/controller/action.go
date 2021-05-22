package controller

import (
	"github.com/kyleu/admini/app/project/action"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/views"
	"github.com/pkg/errors"
	"net/http"
	"strings"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vproject"

	"github.com/gorilla/mux"
)

func ActionWorkbench(w http.ResponseWriter, r *http.Request) {
	act("action.workbench", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		view, err := as.Projects.LoadView(key)
		if err != nil {
			return "", errors.Wrap(err, "unable to load project ["+key+"]")
		}
		ps.Title = view.Project.Name() + " Actions"
		ps.Data = view.Project.Actions
		return render(r, w, as, &vproject.ActionWorkbench{View: view}, ps, "projects", view.Project.Key, "Actions")
	})
}

func ActionOrdering(w http.ResponseWriter, r *http.Request) {
	act("action.ordering", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		prj, err := as.Projects.Load(key, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to load project ["+key+"]")
		}

		_ = r.ParseForm()
		ordering := r.Form.Get("ordering")
		actOrders := action.Orderings{}
		err = util.FromJSON([]byte(ordering), &actOrders)
		if err != nil {
			return "", err
		}
		newActs, err := action.ReorderActions(prj.Actions, actOrders)
		if err != nil {
			return "", err
		}

		err = action.Save(prj.Key, newActs, currentApp.Files)
		if err != nil {
			return "", err
		}

		_, err = as.Projects.Load(prj.Key, true)
		if err != nil {
			return "", err
		}

		return flashAndRedir(true, "updated actions", as.Route("action.workbench", "key", key), w, r, ps)
	})
}

func ActionEdit(w http.ResponseWriter, r *http.Request) {
	act("action.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		prj, err := as.Projects.Load(key, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to load project ["+key+"]")
		}

		pkgString := r.URL.Path
		pkgIdx := strings.Index(pkgString, "/action")
		pkgString = pkgString[pkgIdx + 7:]
		pkg := util.SplitAndTrim(pkgString, "/")

		a, _ := prj.Actions.Get(pkg)
		if a == nil {
			return "", errors.New("no action available at [" + pkgString + "]")
		}

		ps.Title = a.Name()
		ps.Data = a
		return render(r, w, as, &views.TODO{Message: "TODO: edit action [" + pkgString + "]"}, ps, "projects", prj.Key)
	})
}
