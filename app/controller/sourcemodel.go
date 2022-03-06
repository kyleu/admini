package controller

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"admini.dev/app"
	"admini.dev/app/controller/cutil"
	"admini.dev/app/lib/schema"
	"admini.dev/app/lib/schema/model"
	"admini.dev/app/source"
	"admini.dev/app/util"
	"admini.dev/views/vsource"
)

func SourceModelDetail(rc *fasthttp.RequestCtx) {
	act("source.model.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		src, sch, m, err := loadSourceModel(rc, as)
		if err != nil {
			return "", errors.Wrap(err, "")
		}

		ps.Title = src.Name()
		ps.Data = util.ValueMap{sourceKey: src, "schema": sch}
		page := &vsource.ModelDetail{Source: src, Schema: sch, Model: m}
		return render(rc, as, page, ps, "sources", src.Key)
	})
}

func SourceModelSave(rc *fasthttp.RequestCtx) {
	act("source.model.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		src, _, m, err := loadSourceModel(rc, as)
		if err != nil {
			return "", err
		}

		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		m, overrides, err := applyOverrides(frm, m)
		if err != nil {
			return "", errors.Wrap(err, "unable to process overrides")
		}

		if len(overrides) > 0 {
			curr, err := as.Services.Sources.GetOverrides(src.Key)
			if err != nil {
				return "", errors.Wrap(err, "unable to load current overrides")
			}

			final := append(schema.Overrides{}, curr.Purge(m.Path())...)
			final = append(final, overrides...)

			err = as.Services.Sources.SaveOverrides(src.Key, final)
			if err != nil {
				return "", errors.Wrapf(err, "unable to save [%d] overrides", len(curr))
			}
		}

		msg := fmt.Sprintf("saved model [%s] with [%d] overrides", m.Name(), len(overrides))
		return flashAndRedir(true, msg, fmt.Sprintf("/source/%s", src.Key), rc, ps)
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
			if s == util.StringToSingular(util.StringToTitle(m.Key)) {
				m.Title = ""
			} else {
				ret = append(ret, schema.NewOverride("model", m.Path(), schema.KeyTitle, v))
				m.Title = s
			}
		case k == schema.KeyPlural:
			if s == util.StringToPlural(util.StringToTitle(m.Key)) {
				m.Plural = ""
			} else {
				ret = append(ret, schema.NewOverride("model", m.Path(), schema.KeyPlural, v))
				m.Plural = s
			}
		case strings.HasPrefix(k, "f."):
			fs, prop := util.StringSplitLast(k, '.', true)
			fs = strings.TrimPrefix(fs, "f.")
			_, f := m.Fields.Get(fs)
			if f == nil {
				return nil, nil, errors.Errorf("no field found matching [%s]", fs)
			}
			switch prop {
			case schema.KeyTitle:
				if s == util.StringToSingular(util.StringToTitle(f.Key)) {
					f.Title = ""
				} else {
					ret = append(ret, schema.NewOverride("field", append(m.Path(), f.Key), schema.KeyTitle, v))
					f.Title = s
				}
			case schema.KeyPlural:
				if s == util.StringToPlural(util.StringToTitle(f.Key)) {
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

func loadSourceModel(rc *fasthttp.RequestCtx, as *app.State) (*source.Source, *schema.Schema, *model.Model, error) {
	key, err := RCRequiredString(rc, "key", false)
	if err != nil {
		return nil, nil, nil, err
	}
	src, err := as.Services.Sources.Load(key, false)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "unable to load source [%s]", key)
	}
	sch, err := as.Services.Sources.LoadSchema(key)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "unable to load schema for source [%s]", key)
	}
	path, err := RCRequiredString(rc, "path", false)
	if err != nil {
		return nil, nil, nil, err
	}
	pkg := util.Pkg(util.StringSplitAndTrim(path, "/"))

	m := sch.Models.Get(pkg.Drop(1), pkg.Last())
	if m == nil {
		return nil, nil, nil, errors.Errorf("no model found at path [%s]", pkg.ToPath())
	}
	return src, sch, m, nil
}
