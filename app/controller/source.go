package controller

import (
	"fmt"

	"github.com/kyleu/admini/app/util"
	"github.com/valyala/fasthttp"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vsource"
)

const sourceKey = "source"

func SourceList(ctx *fasthttp.RequestCtx) {
	act("source.list", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		s, err := as.Sources.List()
		if err != nil {
			return "", errors.Wrap(err, "unable to load source list")
		}
		ps.Title = "Sources"
		ps.Data = s
		return render(ctx, as, &vsource.List{Sources: s}, ps, "sources")
	})
}

func SourceDetail(ctx *fasthttp.RequestCtx) {
	act("source.detail", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		src, err := as.Sources.Load(key, false)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load source [%s]", key)
		}
		sch, err := as.Sources.LoadSchema(key)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load schema for source [%s]", key)
		}
		ps.Title = src.Name()
		ps.Data = util.ValueMap{sourceKey: src, "schema": sch}
		return render(ctx, as, &vsource.Detail{Source: src, Schema: sch}, ps, "sources", src.Key)
	})
}

func SourceRefresh(ctx *fasthttp.RequestCtx) {
	act("source.refresh", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		_, elapsedMillis, err := as.Sources.SchemaRefresh(key)
		if err != nil {
			return "", errors.Wrapf(err, "unable to  refresh schema for source [%s]", key)
		}

		msg := fmt.Sprintf("refreshed in [%.3fms]", elapsedMillis)
		return flashAndRedir(true, msg, fmt.Sprintf("/source/%s", key), ctx, ps)
	})
}

func SourceModelDetail(ctx *fasthttp.RequestCtx) {
	act("source.model.detail", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		src, err := as.Sources.Load(key, false)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load source [%s]", key)
		}
		sch, err := as.Sources.LoadSchema(key)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load schema for source [%s]", key)
		}
		path, err := ctxRequiredString(ctx, "path", false)
		if err != nil {
			return "", err
		}
		pkg := util.Pkg(util.SplitAndTrim(path, "/"))

		m := sch.Models.Get(pkg.Drop(1), pkg.Last())
		if m == nil {
			return "", errors.Errorf("no model found at path [%s]", pkg.ToPath())
		}

		ps.Title = src.Name()
		ps.Data = util.ValueMap{sourceKey: src, "schema": sch}
		page := &vsource.ModelDetail{Source: src, Schema: sch, Model: m}
		return render(ctx, as, page, ps, "sources", src.Key)
	})
}

func SourceModelSave(ctx *fasthttp.RequestCtx) {
	act("source.model.save", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := ctxRequiredString(ctx, "key", false)
		if err != nil {
			return "", err
		}
		msg := "saved!"
		return flashAndRedir(true, msg, fmt.Sprintf("/source/%s", key), ctx, ps)
	})
}
