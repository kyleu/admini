package postgres

import (
	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/model"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/queries"
)

type tableResult struct {
	Schema string `db:"schema"`
	Name   string `db:"name"`
	Type   string `db:"type"`
	Owner  string `db:"owner"`
}

func (t tableResult) ToModel(logger *zap.SugaredLogger) *model.Model {
	ret := &model.Model{
		Key: t.Name,
		Pkg: util.Pkg{t.Schema},
	}

	switch t.Type {
	case "table":
		ret.Type = model.ModelTypeStruct
	case "view":
		ret.Type = model.ModelTypeInterface
	case "sequence":
		ret.Type = model.ModelTypeSequence
	default:
		logger.Warn("unknown model type [" + t.Type + "]")
		ret.Type = model.ModelTypeUnknown
	}
	return ret
}

func loadTables(enums model.Models, db *database.Service, logger *zap.SugaredLogger) (model.Models, error) {
	tables := []*tableResult{}
	err := db.Select(&tables, queries.ListTables(db.SchemaName), nil)
	if err != nil {
		return nil, errors.Wrap(err, "can't list tables")
	}

	logger.Infof("loading [%v] tables and [%v] enums", len(tables), len(enums))

	ret := make(model.Models, 0, len(tables))
	for _, t := range tables {
		ret = append(ret, t.ToModel(logger))
	}
	return ret, nil
}
