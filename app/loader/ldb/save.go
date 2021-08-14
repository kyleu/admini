package ldb

import (
	"context"
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Save(ctx context.Context, db *database.Service, m *model.Model, ids []interface{}, changes util.ValueMap, logger *zap.SugaredLogger) ([]interface{}, error) {
	cols, vals := changes.KeysAndValues()

	pk := m.GetPK(logger)

	where := make([]string, 0, len(pk))
	for idx, x := range pk {
		where = append(where, fmt.Sprintf("%s = $%d", x, len(vals)+idx+1))
	}

	q := database.SQLUpdateReturning(m.Key, cols, strings.Join(where, " and "), pk)
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
