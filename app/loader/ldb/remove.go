package ldb

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Remove(db *database.Service, m *model.Model, fields []string, values []interface{}, expected int, logger *zap.SugaredLogger) (int, error) {
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
		where.WriteString(fmt.Sprintf(`"%s" = $%d`, x, idx+1))
	}
	q := database.SQLDelete(m.Path().Quoted(), where.String())

	rowsAffected, err := db.Delete(q, nil, expected, values...)
	if err != nil {
		return 0, errors.Wrap(err, "error deleting rows")
	}

	return rowsAffected, nil
}
