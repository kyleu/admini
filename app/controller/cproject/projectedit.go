package cproject

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/project"
	"admini.dev/admini/views/vproject"
)

func ProjectNew(w http.ResponseWriter, r *http.Request) {
	controller.Act("project.new", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "New Project"
		p := &project.Project{}
		ps.Data = p
		avail, err := as.Services.Sources.List(ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to list sources")
		}
		return controller.Render(w, r, as, &vproject.New{Project: p, AvailableSources: avail}, ps, "projects", "New")
	})
}

func ProjectInsert(w http.ResponseWriter, r *http.Request) {
	controller.Act("project.insert", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(r, ps.RequestBody)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}
		key, err := frm.GetString("key", false)
		if err != nil {
			return controller.FlashAndRedir(false, err.Error(), "/project/_new", w, ps)
		}
		title := frm.GetStringOpt("title")
		icon := frm.GetStringOpt("icon")
		description := frm.GetStringOpt("description")
		sources, err := frm.GetStringArray("sources", true)
		if err != nil {
			return "", err
		}
		ret := &project.Project{Key: key, Title: title, Icon: icon, Description: description, Sources: sources}
		err = as.Services.Projects.Save(ret, false, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to save project")
		}
		return controller.FlashAndRedir(true, "saved new project", fmt.Sprintf("/project/%s", key), w, ps)
	})
}

func ProjectEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("project.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.PathString(r, "key", false)
		if err != nil {
			return "", err
		}
		prj, err := as.Services.Projects.LoadRequired(key, false, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load project [%s]", key)
		}
		ps.Title = fmt.Sprintf("Edit [%s]", prj.Name())
		ps.Data = prj

		avail, err := as.Services.Sources.List(ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to list sources")
		}
		return controller.Render(w, r, as, &vproject.Edit{Project: prj, AvailableSources: avail}, ps, "projects", prj.Key, "Edit")
	})
}

func ProjectSave(w http.ResponseWriter, r *http.Request) {
	controller.Act("project.save", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(r, ps.RequestBody)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		key, err := cutil.PathString(r, "key", false)
		if err != nil {
			return "", err
		}

		prj, err := as.Services.Projects.LoadRequired(key, false, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load project [%s]", key)
		}

		prj.Title = frm.GetStringOpt("title")
		prj.Icon = frm.GetStringOpt("icon")
		prj.Description = frm.GetStringOpt("description")
		prj.Sources, err = frm.GetStringArray("sources", true)
		if err != nil {
			return "", err
		}

		err = as.Services.Projects.Save(prj, true, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to save project [%s]", key)
		}

		msg := fmt.Sprintf(`saved project %q`, key)
		return controller.FlashAndRedir(true, msg, fmt.Sprintf("/project/%s", key), w, ps)
	})
}

func ProjectDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("project.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.PathString(r, "key", false)
		if err != nil {
			return "", err
		}
		err = as.Services.Projects.Delete(key, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete project [%s]", key)
		}

		msg := fmt.Sprintf(`deleted project %q`, key)
		return controller.FlashAndRedir(true, msg, "/project", w, ps)
	})
}
