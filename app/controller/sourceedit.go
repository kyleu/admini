package controller

import (
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vsource"
)

func SourceNew(ctx *fasthttp.RequestCtx) {
	act("source.new", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "New Source"
		t := schema.OriginPostgres
		s := &source.Source{Type: t}
		ps.Data = s
		return render(ctx, as, &vsource.New{Origin: t}, ps, "sources", "New")
	})
}

func SourceInsert(ctx *fasthttp.RequestCtx) {
	act("source.insert", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(ctx)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		key, err := frm.GetString("key", false)
		if err != nil {
			return flashError(err, "/source/_new", ctx, ps)
		}
		title, err := frm.GetString("title", true)
		if err != nil {
			return "", err
		}
		description, err := frm.GetString("description", true)
		if err != nil {
			return "", err
		}
		typ, err := frm.GetString("type", true)
		if err != nil {
			return "", err
		}
		ret := currentApp.Sources.NewSource(key, title, description, schema.OriginFromString(typ))
		err = currentApp.Sources.Save(ret, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to save source")
		}
		return flashAndRedir(true, "saved new source", fmt.Sprintf("/source/%s", key), ctx, ps)
	})
}

func SourceEdit(ctx *fasthttp.RequestCtx) {
	act("source.edit", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		src, err := as.Sources.Load(key, false)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load source [%s]", key)
		}
		ps.Title = fmt.Sprintf("Edit [%s]", src.Name())
		ps.Data = src

		switch src.Type {
		case schema.OriginPostgres:
			pcfg := &database.DBParams{}
			if len(src.Config) > 0 {
				err = util.FromJSON(src.Config, pcfg)
				if err != nil {
					return "", errors.Wrap(err, "can't parse postgres config")
				}
			}
			return render(ctx, as, &vsource.EditPostgres{Source: src, Cfg: pcfg}, ps, "sources", src.Key, "Edit")
		default:
			msg := fmt.Sprintf("unhandled source type [%s]", src.Type.String())
			return flashAndRedir(false, msg, "/source", ctx, ps)
		}
	})
}

func SourceSave(ctx *fasthttp.RequestCtx) {
	act("source.save", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(ctx)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		src, err := as.Sources.Load(key, false)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load source [%s]", key)
		}

		src.Title, err = frm.GetString("title", true)
		if err != nil {
			return "", err
		}
		src.Description, err = frm.GetString("description", true)
		if err != nil {
			return "", err
		}

		switch src.Type {
		case schema.OriginPostgres:
			ps, _ := frm.GetString("port", true)
			params := &database.DBParams{}
			params.Host, err = frm.GetString("host", false)
			if err != nil {
				return "", err
			}
			if ps != "" {
				params.Port, _ = strconv.Atoi(ps)
			}
			params.Username, _ = frm.GetString("username", true)
			params.Password, _ = frm.GetString("password", true)
			params.Database, _ = frm.GetString("database", true)
			params.Schema, _ = frm.GetString("schema", true)

			src.Config = util.ToJSONBytes(params, true)
		default:
			return "", errors.Errorf("unable to parse config for source type [%s]", src.Type.String())
		}

		err = currentApp.Sources.Save(src, true)
		if err != nil {
			return "", errors.Wrapf(err, "unable to save source [%s]", key)
		}

		msg := fmt.Sprintf(`saved source "%s"`, key)
		return flashAndRedir(true, msg, fmt.Sprintf("/source/%s", key), ctx, ps)
	})
}

func SourceDelete(ctx *fasthttp.RequestCtx) {
	act("source.delete", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		err = as.Sources.Delete(key)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete source [%s]", key)
		}

		msg := fmt.Sprintf(`deleted source "%s"`, key)
		return flashAndRedir(true, msg, "/source", ctx, ps)
	})
}
