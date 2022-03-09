package mysql

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/util"
	"admini.dev/admini/queries/qmysql"
)

type tableResult struct {
	Schema  string `db:"TABLE_SCHEMA"`
	Name    string `db:"TABLE_NAME"`
	Rows    int64  `db:"TABLE_ROWS"`
	Type    string `db:"TABLE_TYPE"`
	Comment string `db:"TABLE_COMMENT"`
}

func (t tableResult) ToModel(logger *zap.SugaredLogger) *model.Model {
	ret := model.NewModel(util.Pkg{t.Schema}, t.Name)

	switch t.Type {
	case "table", "":
		ret.Type = model.TypeStruct
	case "view":
		ret.Type = model.TypeInterface
	default:
		logger.Warnf("unknown model type [%s] for column [%s]", t.Type, t.Name)
		ret.Type = model.TypeUnknown
	}
	return ret
}

func loadTables(ctx context.Context, enums model.Models, db *database.Service, logger *zap.SugaredLogger) (model.Models, error) {
	var tables []*tableResult
	err := db.Select(ctx, &tables, qmysql.ListTables(db.SchemaName), nil)
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
