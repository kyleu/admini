package ldb

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strings"

	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/util"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/result"
)

var publicSchema = "public"

func List(db *database.Service, m *model.Model, params util.ParamSet, logger *zap.SugaredLogger) (*result.Result, error) {
	p := params.Get(m.Key, m.Fields.Names(), logger)
	if p != nil && p.Limit == 0 {
		p.Limit = 100
	}
	q := modelListQuery(m, p)
	rows, err := db.Query(q, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error listing models for [%s]", m.Key)
	}

	count, err := Count(db, m)
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

func Count(db *database.Service, m *model.Model) (int, error) {
	q := modelCountQuery(m)
	c := struct {
		C int `db:"c"`
	}{}
	if err := db.Get(&c, q, nil); err != nil {
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

func forTable(m *model.Model) (string, string) {
	cols := make([]string, 0, len(m.Fields))
	for _, f := range m.Fields {
		cols = append(cols, fmt.Sprintf(`"%s"`, f.Key))
	}
	tbl := "\"" + m.Key + "\""
	if len(m.Pkg) > 0 {
		l := m.Pkg.Last()
		if l != publicSchema {
			tbl = "\"" + l + "\"." + tbl
		}
	}
	return strings.Join(cols, ", "), tbl
}
