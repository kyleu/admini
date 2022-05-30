package controller

import (
	"context"
	"fmt"

	"admini.dev/admini/app/lib/menu"
	"admini.dev/admini/app/util"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/project"
	"admini.dev/admini/views/vproject"
)

func ProjectNew(rc *fasthttp.RequestCtx) {
	act("project.new", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "New Project"
		p := &project.Project{}
		ps.Data = p
		avail, err := as.Services.Sources.List(ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to list sources")
		}
		return render(rc, as, &vproject.New{Project: p, AvailableSources: avail}, ps, "projects", "New")
	})
}

func ProjectInsert(rc *fasthttp.RequestCtx) {
	act("project.insert", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}
		key, err := frm.GetString("key", false)
		if err != nil {
			return flashAndRedir(false, err.Error(), "/project/_new", rc, ps)
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
		return flashAndRedir(true, "saved new project", fmt.Sprintf("/project/%s", key), rc, ps)
	})
}

func ProjectEdit(rc *fasthttp.RequestCtx) {
	act("project.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", false)
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
		return render(rc, as, &vproject.Edit{Project: prj, AvailableSources: avail}, ps, "projects", prj.Key, "Edit")
	})
}

func ProjectSave(rc *fasthttp.RequestCtx) {
	act("project.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		key, err := cutil.RCRequiredString(rc, "key", false)
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
		return flashAndRedir(true, msg, fmt.Sprintf("/project/%s", key), rc, ps)
	})
}

func ProjectDelete(rc *fasthttp.RequestCtx) {
	act("project.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		err = as.Services.Projects.Delete(key, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete project [%s]", key)
		}

		msg := fmt.Sprintf(`deleted project %q`, key)
		return flashAndRedir(true, msg, "/project", rc, ps)
	})
}

func projectItems(ctx context.Context, as *app.State, logger util.Logger) menu.Items {
	ps, err := as.Services.Projects.List(ctx, logger)
	if err != nil {
		return menu.Items{{Key: "error", Title: "Error", Description: err.Error()}}
	}

	prjMenu := make(menu.Items, 0, len(ps))
	for _, p := range ps {
		prjMenu = append(prjMenu, &menu.Item{
			Key:         p.Key,
			Title:       p.Name(),
			Icon:        p.IconWithFallback(),
			Description: p.Description,
			Route:       "/project/" + p.Key,
		})
	}
	return prjMenu
}
