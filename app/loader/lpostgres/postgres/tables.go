package postgres

import (
	"context"

	model2 "github.com/kyleu/admini/app/schema/model"
	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/queries/qpostgres"
)

type tableResult struct {
	Schema string `db:"schema"`
	Name   string `db:"name"`
	Type   string `db:"type"`
	Owner  string `db:"owner"`
}

func (t tableResult) ToModel(logger *zap.SugaredLogger) *model2.Model {
	ret := model2.NewModel(util.Pkg{t.Schema}, t.Name)

	switch t.Type {
	case "table":
		ret.Type = model2.TypeStruct
	case "view":
		ret.Type = model2.TypeInterface
	default:
		logger.Warnf("unknown model type [%s]", t.Type)
		ret.Type = model2.TypeUnknown
	}
	return ret
}

func loadTables(ctx context.Context, enums model2.Models, db *database.Service, logger *zap.SugaredLogger) (model2.Models, error) {
	var tables []*tableResult
	err := db.Select(ctx, &tables, qpostgres.ListTables(db.SchemaName), nil)
	if err != nil {
		return nil, errors.Wrap(err, "can't list tables")
	}

	logger.Infof("loading [%d] tables and [%d] enums", len(tables), len(enums))

	ret := make(model2.Models, 0, len(tables))
	for _, t := range tables {
		ret = append(ret, t.ToModel(logger))
	}
	return ret, nil
}
