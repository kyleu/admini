package ldb

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema/model"
)

func Remove(
	ctx context.Context, db *database.Service, m *model.Model, fields []string, values []any, expected int, logger *zap.SugaredLogger,
) (int, error) {
	if len(fields) == 0 {
		return 0, errors.New("must provide at least one column")
	}
	if len(fields) != len(values) {
		return 0, errors.Errorf("mismatched lengths between columns (%d) and values (%d)", len(fields), len(values))
	}
	where := strings.Builder{}
	for idx, x := range fields {
		if idx > 0 {
			where.WriteString(" and ")
		}
		if db.Type.Placeholder == "?" {
			where.WriteString(fmt.Sprintf(`%s%s%s = ?`, db.Type.Quote, x, db.Type.Quote))
		} else {
			where.WriteString(fmt.Sprintf(`%s%s%s = $%d`, db.Type.Quote, x, db.Type.Quote, idx+1))
		}
	}
	q := database.SQLDelete(m.Path().Quoted(db.Type.Quote), where.String())

	rowsAffected, err := db.Delete(ctx, q, nil, expected, logger, values...)
	if err != nil {
		return 0, errors.Wrap(err, "error deleting rows")
	}

	return rowsAffected, nil
}
