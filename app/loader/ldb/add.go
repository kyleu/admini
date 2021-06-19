package ldb

import (
	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Add(db *database.Service, m *model.Model, changes util.ValueMap, logger *zap.SugaredLogger) ([]interface{}, error) {
	columns := make([]string, 0, len(changes))
	data := make([]interface{}, 0, len(changes))

	for _, f := range m.Fields {
		v, ok := changes[f.Key]
		if ok {
			columns = append(columns, f.Key)
			data = append(data, v)
		}
	}

	q := database.SQLInsert(m.Path().Quoted(), columns, 1)
	err := db.Insert(q, nil, data...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to insert row")
	}

	return data, nil
}


