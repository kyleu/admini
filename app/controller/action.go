package controller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/views"
	"github.com/kyleu/admini/views/vaction"

	"github.com/kyleu/admini/app/project/action"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app"
)

func ActionOrdering(w http.ResponseWriter, r *http.Request) {
	act("action.ordering", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		prj, err := as.Projects.LoadRequired(key, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to load project ["+key+"]")
		}

		frm, err := cutil.ParseForm(r)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}
		ordering, err := frm.GetString("ordering", true)
		if err != nil {
			return "", err
		}

		actOrders := action.Orderings{}
		err = util.FromJSON([]byte(ordering), &actOrders)
		if err != nil {
			return "", err
		}
		startNanos := time.Now().UnixNano()
		newActs, err := action.ReorderActions(prj.Actions, actOrders)
		if err != nil {
			return "", err
		}

		count, err := action.Save(prj.Key, newActs, currentApp.Files)
		if err != nil {
			return "", err
		}
		elapsedMillis := float64((time.Now().UnixNano()-startNanos)/int64(time.Microsecond)) / float64(1000)

		_, err = as.Projects.LoadRequired(prj.Key, true)
		if err != nil {
			return "", err
		}
		msg := fmt.Sprintf("saved [%v] actions in [%.3fms]", count, elapsedMillis)
		return flashAndRedir(true, msg, as.Route("project.detail", "key", key), w, r, ps)
	})
}

func ActionEdit(w http.ResponseWriter, r *http.Request) {
	act("action.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		v, a, err := loadAction(r, as)
		if err != nil {
			return "", errors.Wrap(err, "error loading project and action")
		}
		ps.Title = a.Name()
		ps.Data = a
		page := &vaction.Edit{View: v, Act: a}
		return render(r, w, as, page, ps, append([]string{"projects", v.Project.Key}, a.Path()...)...)
	})
}

func ActionSave(w http.ResponseWriter, r *http.Request) {
	act("action.save", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		v, a, err := loadAction(r, as)
		if err != nil {
			return "", errors.Wrap(err, "error loading project and action")
		}
		ps.Title = a.Name()
		ps.Data = a
		page := &views.TODO{Message: "Nice work!"}
		return render(r, w, as, page, ps, append([]string{"projects", v.Project.Key}, a.Path()...)...)
	})
}

func loadAction(r *http.Request, as *app.State) (*project.View, *action.Action, error) {
	key := mux.Vars(r)["key"]
	v, err := as.Projects.LoadView(key)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to load project ["+key+"]")
	}

	pkgString := r.URL.Path
	pkgIdx := strings.Index(pkgString, "/action")
	pkgString = pkgString[pkgIdx+7:]
	pkg := util.SplitAndTrim(pkgString, "/")

	a, _ := v.Project.Actions.Get(pkg)
	if a == nil {
		return nil, nil, errors.New("no action available at [" + pkgString + "]")
	}

	return v, a, nil
}
