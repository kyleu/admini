package ldb

import (
	"context"
	"strings"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/filter"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/result"
	"github.com/pkg/errors"
)

func List(ctx context.Context, db *database.Service, m *model.Model, opts *filter.Options) (*result.Result, error) {
	p := opts.Params
	if p != nil && p.Limit == 0 {
		p.Limit = filter.MaxRowsDefault
	}
	q := modelListQuery(db.Type, m, p)
	rows, err := db.Query(ctx, q, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error listing models for [%s]", m.String())
	}
	defer func() { _ = rows.Close() }()

	count, err := Count(ctx, db, m)
	if err != nil {
		return nil, errors.Wrapf(err, "error constructing result for [%s]", m.String())
	}

	var timing *result.Timing
	ret, err := ParseResultFields(m.Name(), count, q, timing, m.Fields, rows)
	if err != nil {
		return nil, errors.Wrapf(err, "error constructing result for [%s]", m.String())
	}

	return ret, nil
}

func Count(ctx context.Context, db *database.Service, m *model.Model) (int, error) {
	q := modelCountQuery(db.Type, m)
	c := struct {
		C int `db:"c"`
	}{}
	if err := db.Get(ctx, &c, q, nil); err != nil {
		return 0, errors.Wrapf(err, "error listing models for [%s]", m.String())
	}
	return c.C, nil
}

func modelListQuery(typ *database.DBType, m *model.Model, params *filter.Params) string {
	cols, tbl := forTable(typ, m)
	return database.SQLSelect(cols, tbl, "", params.OrderByString(), params.Limit, params.Offset)
}

func modelCountQuery(typ *database.DBType, m *model.Model) string {
	return database.SQLSelectSimple("count(*) as c", m.Path().Quoted(typ.Quote), "")
}

func forTable(typ *database.DBType, m *model.Model) (string, string) {
	cols := make([]string, 0, len(m.Fields))
	for _, f := range m.Fields {
		cols = append(cols, typ.Quote + f.Key + typ.Quote)
	}
	return strings.Join(cols, ", "), m.Path().Quoted(typ.Quote)
}
