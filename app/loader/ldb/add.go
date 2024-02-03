package ldb

import (
	"context"

	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/util"
)

func Add(ctx context.Context, db *database.Service, m *model.Model, changes util.ValueMap, logger util.Logger) ([]any, error) {
	columns, data := changes.KeysAndValues()
	pk := m.GetPK(logger)

	if db.Type.SupportsReturning {
		q := database.SQLInsertReturning(m.Path().Quoted(db.Type.Quote), columns, 1, pk, db.Type)
		out, err := db.QuerySingleRow(ctx, q, nil, logger, data...)
		if err != nil {
			return nil, errors.Wrap(err, "unable to insert row reurning primary key")
		}
		ret := make([]any, 0, len(out))
		for _, p := range pk {
			ret = append(ret, out[p])
		}
		return ret, nil
	}

	tx, err := db.StartTransaction(logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open transaction")
	}
	iq := database.SQLInsert(m.Path().Quoted(db.Type.Quote), columns, 1, db.Type)
	_, err = db.Exec(ctx, iq, tx, 1, logger, data...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to insert row")
	}

	pkVals := make([]any, 0, len(pk))
	for _, x := range pk {
		hit := false
		for idx, c := range columns {
			if c == x {
				hit = true
				pkVals = append(pkVals, data[idx])
				break
			}
		}
		if !hit {
			return nil, errors.Errorf("missing pk field [%s] for insert", x)
		}
	}

	return loadAfterEdit(ctx, pk, pkVals, m, tx, db, logger)
}
