package ldb

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/kyleu/admini/app/lib/database"
	"github.com/kyleu/admini/app/lib/schema/model"
	"github.com/kyleu/admini/app/util"
)

func Save(
	ctx context.Context, db *database.Service, m *model.Model, ids []interface{}, changes util.ValueMap, logger *zap.SugaredLogger,
) ([]interface{}, error) {
	cols, vals := changes.KeysAndValues()

	pk := m.GetPK(logger)

	where := make([]string, 0, len(pk))
	for idx, x := range pk {
		if db.Type.Placeholder == "?" {
			where = append(where, fmt.Sprintf("%s = ?", x))
		} else {
			where = append(where, fmt.Sprintf("%s = $%d", x, len(vals)+idx+1))
		}
	}

	if db.Type.SupportsReturning {
		q := database.SQLUpdateReturning(m.Key, cols, strings.Join(where, " and "), pk, db.Type.Placeholder)
		out, err := db.QuerySingleRow(ctx, q, nil, append(vals, ids...)...)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to save [%s] with primary key [%s]", m.Name(), strings.Join(pk, "::"))
		}
		ret := make([]interface{}, 0, len(out))
		for _, p := range pk {
			ret = append(ret, out[p])
		}
		return ret, nil
	}

	tx, err := db.StartTransaction()
	if err != nil {
		return nil, err
	}

	uq := database.SQLUpdate(m.Key, cols, strings.Join(where, " and "), db.Type.Placeholder)
	_, err = db.Exec(ctx, uq, tx, -1, append(vals, ids...)...)
	if err != nil {
		_ = tx.Rollback()
		return nil, errors.Wrap(err, "unable to insert row")
	}

	return loadAfterEdit(ctx, pk, ids, m, tx, db)
}

func loadAfterEdit(ctx context.Context, pk []string, pkVals []interface{}, m *model.Model, tx *sqlx.Tx, db *database.Service) ([]interface{}, error) {
	wc := make([]string, 0, len(pk))
	for idx, x := range pk {
		if db.Type.Placeholder == "?" {
			wc = append(wc, fmt.Sprintf(`%s%s%s = ?`, db.Type.Quote, x, db.Type.Quote))
		} else {
			wc = append(wc, fmt.Sprintf(`%s%s%s = $%d`, db.Type.Quote, x, db.Type.Quote, idx+1))
		}
	}
	sq := database.SQLSelectSimple(strings.Join(pk, ", "), m.Path().Quoted(db.Type.Quote), strings.Join(wc, " and "))
	out, err := db.QuerySingleRow(ctx, sq, tx, pkVals...)
	if err != nil {
		_ = tx.Rollback()
		return nil, errors.Wrap(err, "unable to select newly-inserted row")
	}
	ret := make([]interface{}, 0, len(pk))
	for _, p := range pk {
		ret = append(ret, out[p])
	}
	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "unable to commit insert transaction")
	}

	return ret, nil
}
