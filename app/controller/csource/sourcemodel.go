package csource

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/schema"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/source"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views/vsource"
)

func SourceModelDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("source.model.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		src, sch, m, err := loadSourceModel(r, as, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "")
		}

		ps.Title = src.Name()
		ps.Data = util.ValueMap{sourceKey: src, "schema": sch}
		page := &vsource.ModelDetail{Source: src, Schema: sch, Model: m}
		return controller.Render(r, as, page, ps, "sources", src.Key)
	})
}

func SourceModelSave(w http.ResponseWriter, r *http.Request) {
	controller.Act("source.model.save", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		src, _, m, err := loadSourceModel(r, as, ps.Logger)
		if err != nil {
			return "", err
		}

		frm, err := cutil.ParseForm(r, ps.RequestBody)
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
		return controller.FlashAndRedir(true, msg, fmt.Sprintf("/source/%s", src.Key), ps)
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

func loadSourceModel(r *http.Request, as *app.State, logger util.Logger) (*source.Source, *schema.Schema, *model.Model, error) {
	key, err := cutil.PathString(r, "key", false)
	if err != nil {
		return nil, nil, nil, err
	}
	src, err := as.Services.Sources.Load(key, false, logger)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "unable to load source [%s]", key)
	}
	sch, err := as.Services.Sources.LoadSchema(key)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "unable to load schema for source [%s]", key)
	}
	path, err := cutil.PathString(r, "path", false)
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
