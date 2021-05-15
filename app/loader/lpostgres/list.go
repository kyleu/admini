package lpostgres

import (
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/util"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/result"
)

var publicSchema = "public"

func (l *Loader) List(source string, cfg []byte, m *model.Model, params util.ParamSet) (*result.Result, error) {
	db, err := l.openDatabase(source, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database")
	}

	q := modelListQuery(m, params.Get(m.Key, m.Fields.Names(), l.logger))
	rows, err := db.Query(q, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error listing models for [%v]", m.Key)
	}

	count, err := l.Count(source, cfg, m)
	if err != nil {
		return nil, errors.Wrapf(err, "error constructing result for [%v]", m.Key)
	}

	var timing *result.Timing
	ret, err := ParseResultFields(m.Key, count, q, timing, m.Fields, rows)
	if err != nil {
		return nil, errors.Wrapf(err, "error constructing result for [%v]", m.Key)
	}

	return ret, nil
}

func (l *Loader) Count(source string, cfg []byte, m *model.Model) (int, error) {
	db, err := l.openDatabase(source, cfg)
	if err != nil {
		return 0, errors.Wrap(err, "error opening database")
	}

	q := modelCountQuery(m)
	c := struct {
		C int `db:"c"`
	}{}
	err = db.Get(&c, q, nil)
	if err != nil {
		return 0, errors.Wrapf(err, "error listing models for [%v]", m.Key)
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
