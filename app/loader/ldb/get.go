package ldb

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/result"
	"admini.dev/admini/app/util"
)

func Get(ctx context.Context, db *database.Service, m *model.Model, ids []any, logger util.Logger) (*result.Result, error) {
	q, err := modelGetByPKQuery(db.Type, m, logger)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(ctx, q, nil, logger, ids...)
	if err != nil {
		return nil, errors.Wrapf(err, "error listing models for [%s]", m.String())
	}
	defer func() { _ = rows.Close() }()

	var timing *result.Timing
	ret, err := ParseResultFields(m.Name(), 0, q, timing, m.Fields, rows)
	if err != nil {
		return nil, errors.Wrapf(err, "error constructing result for [%s]", m.String())
	}

	return ret, nil
}

func modelGetByPKQuery(typ *database.DBType, m *model.Model, logger util.Logger) (string, error) {
	cols, tbl := forTable(typ, m)
	pk := m.GetPK(logger)
	if len(pk) == 0 {
		return "", errors.Errorf("no PK for model [%s]", m.String())
	}
	where := make([]string, 0, len(pk))
	for idx, pkf := range pk {
		where = append(where, fmt.Sprintf(`%s = %s`, typ.Quoted(pkf), typ.PlaceholderFor(idx+1)))
	}
	return database.SQLSelectSimple(cols, tbl, typ, strings.Join(where, " and ")), nil
}
