package postgres

import (
	"context"

	"github.com/kyleu/admini/app/lib/schema/model"
	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/lib/database"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/queries/qpostgres"
)

type tableResult struct {
	Schema string `db:"schema"`
	Name   string `db:"name"`
	Type   string `db:"type"`
	Owner  string `db:"owner"`
}

func (t tableResult) ToModel(logger *zap.SugaredLogger) *model.Model {
	ret := model.NewModel(util.Pkg{t.Schema}, t.Name)

	switch t.Type {
	case "table":
		ret.Type = model.TypeStruct
	case "view":
		ret.Type = model.TypeInterface
	default:
		logger.Warnf("unknown model type [%s]", t.Type)
		ret.Type = model.TypeUnknown
	}
	return ret
}

func loadTables(ctx context.Context, enums model.Models, db *database.Service, logger *zap.SugaredLogger) (model.Models, error) {
	var tables []*tableResult
	err := db.Select(ctx, &tables, qpostgres.ListTables(db.SchemaName), nil)
	if err != nil {
		return nil, errors.Wrap(err, "can't list tables")
	}

	logger.Infof("loading [%d] tables and [%d] enums", len(tables), len(enums))

	ret := make(model.Models, 0, len(tables))
	for _, t := range tables {
		ret = append(ret, t.ToModel(logger))
	}
	return ret, nil
}
