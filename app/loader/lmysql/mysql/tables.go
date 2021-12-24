package mysql

import (
	"context"

	model2 "github.com/kyleu/admini/app/schema/model"
	"github.com/kyleu/admini/queries/qmysql"
	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/util"
)

type tableResult struct {
	Schema  string `db:"TABLE_SCHEMA"`
	Name    string `db:"TABLE_NAME"`
	Rows    int64  `db:"TABLE_ROWS"`
	Type    string `db:"TABLE_TYPE"`
	Comment string `db:"TABLE_COMMENT"`
}

func (t tableResult) ToModel(logger *zap.SugaredLogger) *model2.Model {
	ret := model2.NewModel(util.Pkg{t.Schema}, t.Name)

	switch t.Type {
	case "table", "":
		ret.Type = model2.TypeStruct
	case "view":
		ret.Type = model2.TypeInterface
	default:
		logger.Warnf("unknown model type [%s] for column [%s]", t.Type, t.Name)
		ret.Type = model2.TypeUnknown
	}
	return ret
}

func loadTables(ctx context.Context, enums model2.Models, db *database.Service, logger *zap.SugaredLogger) (model2.Models, error) {
	var tables []*tableResult
	err := db.Select(ctx, &tables, qmysql.ListTables(db.SchemaName), nil)
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
