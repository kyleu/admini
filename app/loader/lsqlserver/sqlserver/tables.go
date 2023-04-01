package sqlserver

import (
	"context"

	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/util"
	"admini.dev/admini/queries/qsqlserver"
)

type tableResult struct {
	Schema string `db:"schema_name"`
	Name   string `db:"table_name"`
}

func (t tableResult) ToModel(logger util.Logger) *model.Model {
	ret := model.NewModel(util.Pkg{t.Schema}, t.Name)
	ret.Type = model.TypeStruct
	return ret
}

func loadTables(ctx context.Context, enums model.Models, db *database.Service, logger util.Logger) (model.Models, error) {
	var tables []*tableResult
	err := db.Select(ctx, &tables, qsqlserver.ListTables(db.SchemaName), nil, logger)
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
