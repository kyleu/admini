package controller

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/menu"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/vsource"
)

const sourceKey = "source"

func SourceList(rc *fasthttp.RequestCtx) {
	act("source.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		s, err := as.Services.Sources.List(ps.Logger)
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
		return render(rc, as, &vsource.Detail{Source: src, Schema: sch}, ps, "sources", src.Key)
	})
}

func SourceRefresh(rc *fasthttp.RequestCtx) {
	act("source.refresh", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		_, elapsedMillis, err := as.Services.Sources.SchemaRefresh(ps.Context, key, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to  refresh schema for source [%s]", key)
		}

		msg := fmt.Sprintf("refreshed in [%.3fms]", elapsedMillis)
		return flashAndRedir(true, msg, fmt.Sprintf("/source/%s", key), rc, ps)
	})
}

func SourceHack(rc *fasthttp.RequestCtx) {
	act("source.hack", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		sch, err := as.Services.Sources.LoadSchema(key)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load schema for source [%s]", key)
		}
		ret, err := sch.Hack(ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to run schema hack for source [%s]", key)
		}
		ps.Data = ret
		return render(rc, as, &vsource.Hack{Schema: sch, Result: ret}, ps, "sources", key, "hack")
	})
}

func sourceItems(ctx context.Context, as *app.State, logger util.Logger) menu.Items {
	ss, err := as.Services.Sources.List(logger)
	if err != nil {
		return menu.Items{{Key: "error", Title: "Error", Description: err.Error()}}
	}

	srcMenu := make(menu.Items, 0, len(ss))
	for _, s := range ss {
		srcMenu = append(srcMenu, &menu.Item{
			Key:         s.Key,
			Title:       s.Name(),
			Icon:        s.IconWithFallback(),
			Description: s.Description,
			Route:       "/source/" + s.Key,
		})
	}
	return srcMenu
}
