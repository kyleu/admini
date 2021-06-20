package ldb

import (
	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Add(db *database.Service, m *model.Model, changes util.ValueMap, logger *zap.SugaredLogger) ([]interface{}, error) {
	columns, data := changes.KeysAndValues()
	pk := m.GetPK(logger)
	q := database.SQLInsertReturning(m.Path().Quoted(), columns, 1, pk)
	out, err := db.QuerySingleRow(q, nil, data...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to insert row")
	}

	ret := make([]interface{}, 0, len(out))
	for _, p := range pk {
		ret = append(ret, out[p])
	}

	return ret, nil
}
