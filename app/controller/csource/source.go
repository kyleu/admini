package csource

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/vsource"
)

const sourceKey = "source"

func SourceList(rc *fasthttp.RequestCtx) {
	controller.Act("source.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		s, err := as.Services.Sources.List(ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to load source list")
		}
		ps.Title = "Sources"
		ps.Data = s
		return controller.Render(rc, as, &vsource.List{Sources: s}, ps, "sources")
	})
}

func SourceDetail(rc *fasthttp.RequestCtx) {
	controller.Act("source.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		src, err := as.Services.Sources.Load(key, false, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load source [%s]", key)
		}
		sch, err := as.Services.Sources.LoadSchema(key)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load schema for source [%s]", key)
		}
		ps.Title = src.Name()
		ps.Data = util.ValueMap{sourceKey: src, "schema": sch}
		return controller.Render(rc, as, &vsource.Detail{Source: src, Schema: sch}, ps, "sources", src.Key)
	})
}

func SourceRefresh(rc *fasthttp.RequestCtx) {
	controller.Act("source.refresh", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		_, elapsedMillis, err := as.Services.Sources.SchemaRefresh(ps.Context, key, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to  refresh schema for source [%s]", key)
		}

		msg := fmt.Sprintf("refreshed in [%.3fms]", elapsedMillis)
		return controller.FlashAndRedir(true, msg, fmt.Sprintf("/source/%s", key), rc, ps)
	})
}

func SourceHack(rc *fasthttp.RequestCtx) {
	controller.Act("source.hack", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		sch, err := as.Services.Sources.LoadSchema(key)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load schema for source [%s]", key)
		}
		if string(rc.URI().QueryArgs().Peek("x")) == "svc" {
			ret, err := sch.HackSvc(ps.Logger)
			if err != nil {
				return "", errors.Wrapf(err, "unable to run schema hack for source [%s]", key)
			}
			rc.Response.SetBodyRaw([]byte(ret))
			return "", nil
		}
		ret, err := sch.Hack(ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to run schema hack for source [%s]", key)
		}
		ps.Data = ret
		return controller.Render(rc, as, &vsource.Hack{Schema: sch, Result: ret}, ps, "sources", key, "hack")
	})
}
