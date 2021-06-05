package lpostgres

import (
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/util"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/result"
)

var publicSchema = "public"

func (l *Loader) List(m *model.Model, params util.ParamSet) (*result.Result, error) {
	q := modelListQuery(m, params.Get(m.Key, m.Fields.Names(), l.logger))
	rows, err := l.db.Query(q, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error listing models for [%s]", m.Key)
	}

	count, err := l.Count(m)
	if err != nil {
		return nil, errors.Wrapf(err, "error constructing result for [%s]", m.Key)
	}

	var timing *result.Timing
	ret, err := ParseResultFields(m.Name(), count, q, timing, m.Fields, rows)
	if err != nil {
		return nil, errors.Wrapf(err, "error constructing result for [%s]", m.Key)
	}

	return ret, nil
}

func (l *Loader) Count(m *model.Model) (int, error) {
	q := modelCountQuery(m)
	c := struct {
		C int `db:"c"`
	}{}
	if err := l.db.Get(&c, q, nil); err != nil {
		return 0, errors.Wrapf(err, "error listing models for [%s]", m.Key)
	}
	return c.C, nil
}

func modelListQuery(m *model.Model, params *util.Params) string {
	cols, tbl := forTable(m)
	return database.SQLSelect(cols, tbl, "", params.OrderByString(), params.Limit, params.Offset)
}

func modelCountQuery(m *model.Model) string {
	tbl := "\"" + m.Key + "\""
	if len(m.Pkg) > 0 {
		l := m.Pkg.Last()
		if l != publicSchema {
			tbl = "\"" + l + "\"." + tbl
		}
	}

	return database.SQLSelectSimple("count(*) as c", tbl, "")
}
