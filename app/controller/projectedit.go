package controller

import (
	"fmt"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/views/vproject"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

func ProjectNew(ctx *fasthttp.RequestCtx) {
	act("project.new", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "New Project"
		p := &project.Project{}
		ps.Data = p
		avail, err := currentApp.Sources.List()
		if err != nil {
			return "", errors.Wrap(err, "unable to list sources")
		}
		return render(ctx, as, &vproject.New{Project: p, AvailableSources: avail}, ps, "projects", "New")
	})
}

func ProjectInsert(ctx *fasthttp.RequestCtx) {
	act("project.insert", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(ctx)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}
		key, err := frm.GetString("key", false)
		if err != nil {
			return flashError(err, "/project/_new", ctx, ps)
		}
		title, err := frm.GetString("title", true)
		if err != nil {
			return flashError(err, "/project/_new", ctx, ps)
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
		return flashAndRedir(true, "saved new project", fmt.Sprintf("/project/%s", key), ctx, ps)
	})
}

func ProjectEdit(ctx *fasthttp.RequestCtx) {
	act("project.edit", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		prj, err := as.Projects.LoadRequired(key, false)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load project [%s]", key)
		}
		ps.Title = fmt.Sprintf("Edit [%s]", prj.Name())
		ps.Data = prj

		avail, err := currentApp.Sources.List()
		if err != nil {
			return "", errors.Wrap(err, "unable to list sources")
		}
		return render(ctx, as, &vproject.Edit{Project: prj, AvailableSources: avail}, ps, "projects", prj.Key, "Edit")
	})
}

func ProjectSave(ctx *fasthttp.RequestCtx) {
	act("project.save", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(ctx)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}

		prj, err := as.Projects.LoadRequired(key, false)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load project [%s]", key)
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
			return "", errors.Wrapf(err, "unable to save project [%s]", key)
		}

		msg := fmt.Sprintf(`saved project "%s"`, key)
		return flashAndRedir(true, msg, fmt.Sprintf("/project/%s", key), ctx, ps)
	})
}

func ProjectDelete(ctx *fasthttp.RequestCtx) {
	act("project.delete", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		err = as.Projects.Delete(key)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete project [%s]", key)
		}

		msg := fmt.Sprintf(`deleted project "%s"`, key)
		return flashAndRedir(true, msg, "/project", ctx, ps)
	})
}
