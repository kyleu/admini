package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/views/vproject"
	"github.com/pkg/errors"
)

func ProjectNew(w http.ResponseWriter, r *http.Request) {
	act("project.new", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "New Project"
		p := &project.Project{}
		ps.Data = p
		avail, err := currentApp.Sources.List()
		if err != nil {
			return "", errors.Wrap(err, "unable to list sources")
		}
		return render(r, w, as, &vproject.New{Project: p, AvailableSources: avail}, ps, "projects", "New")
	})
}

func ProjectInsert(w http.ResponseWriter, r *http.Request) {
	act("project.insert", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(r)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}
		key, err := frm.GetString("key", false)
		if err != nil {
			return flashError(err, as.Route("project.new", "key", key), w, r, ps)
		}
		title, err := frm.GetString("title", true)
		if err != nil {
			return flashError(err, as.Route("project.new", "key", key), w, r, ps)
		}
		description, err := frm.GetString("description", true)
		if err != nil {
			return "", err
		}
		sources, err := frm.GetStringArray("sources", true)
		if err != nil {
			return "", err
		}
		ret := &project.Project{Key: key, Title: title, Description: description, Sources: sources}
		err = currentApp.Projects.Save(ret, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to save project")
		}
		return flashAndRedir(true, "saved new project", as.Route("project.detail", "key", key), w, r, ps)
	})
}

func ProjectEdit(w http.ResponseWriter, r *http.Request) {
	act("project.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		prj, err := as.Projects.LoadRequired(key, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to load project ["+key+"]")
		}
		ps.Title = "Edit [" + prj.Name() + "]"
		ps.Data = prj

		avail, err := currentApp.Sources.List()
		if err != nil {
			return "", errors.Wrap(err, "unable to list sources")
		}
		return render(r, w, as, &vproject.Edit{Project: prj, AvailableSources: avail}, ps, "projects", prj.Key, "Edit")
	})
}

func ProjectSave(w http.ResponseWriter, r *http.Request) {
	act("project.save", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(r)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		key := mux.Vars(r)["key"]

		prj, err := as.Projects.LoadRequired(key, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to load project ["+key+"]")
		}

		prj.Title, err = frm.GetString("title", true)
		if err != nil {
			return "", err
		}
		prj.Description, err = frm.GetString("description", true)
		if err != nil {
			return "", err
		}
		prj.Sources, err = frm.GetStringArray("sources", true)
		if err != nil {
			return "", err
		}

		err = currentApp.Projects.Save(prj, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to save project ["+key+"]")
		}

		msg := fmt.Sprintf(`saved project "%v"`, key)
		return flashAndRedir(true, msg, as.Route("project.detail", "key", key), w, r, ps)
	})
}

func ProjectDelete(w http.ResponseWriter, r *http.Request) {
	act("project.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]

		err := as.Projects.Delete(key)
		if err != nil {
			return "", errors.Wrap(err, "unable to delete project ["+key+"]")
		}

		msg := fmt.Sprintf(`deleted project "%v"`, key)
		return flashAndRedir(true, msg, as.Route("project.list"), w, r, ps)
	})
}
