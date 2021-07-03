package controller

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/source"

	"github.com/kyleu/admini/app/util"
	"github.com/valyala/fasthttp"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vsource"
)

func SourceModelDetail(ctx *fasthttp.RequestCtx) {
	act("source.model.detail", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		src, sch, m, err := loadSourceModel(ctx, as)
		if err != nil {
			return "", errors.Wrap(err, "")
		}

		ps.Title = src.Name()
		ps.Data = util.ValueMap{sourceKey: src, "schema": sch}
		page := &vsource.ModelDetail{Source: src, Schema: sch, Model: m}
		return render(ctx, as, page, ps, "sources", src.Key)
	})
}

func SourceModelSave(ctx *fasthttp.RequestCtx) {
	act("source.model.save", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		src, _, m, err := loadSourceModel(ctx, as)
		if err != nil {
			return "", err
		}

		frm, err := cutil.ParseForm(ctx)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		m, todo, err := applyOverrides(frm, m)
		if err != nil {
			return "", errors.Wrap(err, "unable to process overrides")
		}

		if len(todo) > 0 {
			curr, err := currentApp.Sources.GetOverrides(src.Key)
			if err != nil {
				return "", errors.Wrap(err, "unable to load current overrides")
			}

			purged := curr.Purge(m.Path())
			final := append(purged, todo...)

			err = currentApp.Sources.SaveOverrides(src.Key, final)
			if err != nil {
				return "", errors.Wrapf(err, "unable to save [%d] overrides", len(curr))
			}
		}

		msg := fmt.Sprintf("saved model [%s] with [%d] overrides", m.Name(), len(todo))
		return flashAndRedir(true, msg, fmt.Sprintf("/source/%s", src.Key), ctx, ps)
	})
}

func applyOverrides(frm util.ValueMap, m *model.Model) (*model.Model, schema.Overrides, error) {
	var ret schema.Overrides

	for k, v := range frm {
		s, ok := v.(string)
		if !ok {
			return nil, nil, errors.Errorf("value is of type [%T], expected [string]", v)
		}
		switch {
		case k == schema.KeyTitle:
			if s == util.ToSingular(util.ToTitle(m.Key)) {
				m.Title = ""
			} else {
				ret = append(ret, schema.NewOverride("model", m.Path(), schema.KeyTitle, v))
				m.Title = s
			}
		case k == schema.KeyPlural:
			if s == util.ToPlural(util.ToTitle(m.Key)) {
				m.Plural = ""
			} else {
				ret = append(ret, schema.NewOverride("model", m.Path(), schema.KeyPlural, v))
				m.Plural = s
			}
		case strings.HasPrefix(k, "f."):
			fs, prop := util.SplitStringLast(k, '.', true)
			fs = strings.TrimPrefix(fs, "f.")
			_, f := m.Fields.Get(fs)
			if f == nil {
				return nil, nil, errors.Errorf("no field found matching [%s]", fs)
			}
			switch prop {
			case schema.KeyTitle:
				if s == util.ToSingular(util.ToTitle(f.Key)) {
					f.Title = ""
				} else {
					ret = append(ret, schema.NewOverride("field", append(m.Path(), f.Key), schema.KeyTitle, v))
					f.Title = s
				}
			case schema.KeyPlural:
				if s == util.ToPlural(util.ToTitle(f.Key)) {
					f.Plural = ""
				} else {
					ret = append(ret, schema.NewOverride("field", append(m.Path(), f.Key), schema.KeyPlural, v))
					f.Plural = s
				}
			default:
				return nil, nil, errors.Errorf("unhandled override field value [%s]", prop)
			}
		default:
			return nil, nil, errors.Errorf("unhandled override value [%s]", k)
		}
	}
	return m, ret, nil
}

func loadSourceModel(ctx *fasthttp.RequestCtx, as *app.State) (*source.Source, *schema.Schema, *model.Model, error) {
	key, err := ctxRequiredString(ctx, "key", false)
	if err != nil {
		return nil, nil, nil, err
	}
	src, err := as.Sources.Load(key, false)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "unable to load source [%s]", key)
	}
	sch, err := as.Sources.LoadSchema(key)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "unable to load schema for source [%s]", key)
	}
	path, err := ctxRequiredString(ctx, "path", false)
	if err != nil {
		return nil, nil, nil, err
	}
	pkg := util.Pkg(util.SplitAndTrim(path, "/"))

	m := sch.Models.Get(pkg.Drop(1), pkg.Last())
	if m == nil {
		return nil, nil, nil, errors.Errorf("no model found at path [%s]", pkg.ToPath())
	}
	return src, sch, m, nil
}
