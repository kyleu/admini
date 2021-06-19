package ldb

import (
	"fmt"
	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strings"
)

func Save(db *database.Service, m *model.Model, ids []interface{}, changes util.ValueMap, logger *zap.SugaredLogger) ([]interface{}, error) {
	cols, vals := changes.KeysAndValues()

	pk := m.GetPK(logger)

	var where []string
	for idx, x := range pk {
		where = append(where, fmt.Sprintf("%s = $%d", x, len(vals) + idx + 1))
	}

	q := database.SQLUpdate(m.Key, cols, strings.Join(where, " and "))

	println(q)

	_, err := db.Update(q, nil, 1, append(vals, ids...)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to save [%s] with primary key [%s]", m.Name(), strings.Join(pk, "::"))
	}

	return vals, nil
}
