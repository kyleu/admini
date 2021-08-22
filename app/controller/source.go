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

func SourceList(rc *fasthttp.RequestCtx) {
	act("source.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		s, err := as.Services.Sources.List()
		if err != nil {
			return "", errors.Wrap(err, "unable to load source list")
		}
		ps.Title = "Sources"
		ps.Data = s
		return render(rc, as, &vsource.List{Sources: s}, ps, "sources")
	})
}

func SourceDetail(rc *fasthttp.RequestCtx) {
	act("source.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := rcRequiredString(rc, "key", false)
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
		return render(rc, as, &vsource.Detail{Source: src, Schema: sch}, ps, "sources", src.Key)
	})
}

func SourceRefresh(rc *fasthttp.RequestCtx) {
	act("source.refresh", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := rcRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		_, elapsedMillis, err := as.Services.Sources.SchemaRefresh(ps.Context, key)
		if err != nil {
			return "", errors.Wrapf(err, "unable to  refresh schema for source [%s]", key)
		}

		msg := fmt.Sprintf("refreshed in [%.3fms]", elapsedMillis)
		return flashAndRedir(true, msg, fmt.Sprintf("/source/%s", key), rc, ps)
	})
}
