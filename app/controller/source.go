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
		s, err := as.Services.Sources.List()
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
		src, err := as.Services.Sources.Load(key, false)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load source [%s]", key)
		}
		sch, err := as.Services.Sources.LoadSchema(key)
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
		_, elapsedMillis, err := as.Services.Sources.SchemaRefresh(ps.Context, key)
		if err != nil {
			return "", errors.Wrapf(err, "unable to  refresh schema for source [%s]", key)
		}

		msg := fmt.Sprintf("refreshed in [%.3fms]", elapsedMillis)
		return flashAndRedir(true, msg, fmt.Sprintf("/source/%s", key), ctx, ps)
	})
}
