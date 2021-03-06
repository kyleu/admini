package mysql

import (
	"context"

	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/util"
)

func LoadDatabaseSchema(ctx context.Context, db *database.Service, logger util.Logger) (*schema.Schema, error) {
	metadata := loadMetadata(ctx, db)

	tables, err := loadTables(ctx, nil, db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "can't load tables")
	}

	err = loadColumns(ctx, tables, db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "can't load columns")
	}

	err = loadIndexes(ctx, tables, db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "can't load indexes")
	}

	err = loadForeignKeys(ctx, tables, db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "can't load foreign keys")
	}

	models := make(model.Models, 0, len(tables))
	models = append(models, tables...)
	models.Sort()

	ret := &schema.Schema{
		Paths:    []string{"mysql:" + db.DatabaseName},
		Models:   models,
		Metadata: metadata,
	}

	return ret, nil
}
