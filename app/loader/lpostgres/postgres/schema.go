package postgres

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

	scalars, err := loadScalars()
	if err != nil {
		return nil, errors.Wrap(err, "can't load scalars")
	}

	enums, err := loadEnums(ctx, db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "can't load enums")
	}

	tables, err := loadTables(ctx, enums, db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "can't load tables")
	}

	models := make(model.Models, 0, len(tables)+len(enums))
	models = append(models, enums...)
	models = append(models, tables...)
	models.Sort()

	err = loadColumns(ctx, models, db, logger)
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

	err = loadEnumRelations(ctx, enums, tables, db)
	if err != nil {
		return nil, errors.Wrap(err, "can't load foreign keys")
	}

	ret := &schema.Schema{
		Paths:    []string{"postgres:" + db.DatabaseName},
		Scalars:  scalars,
		Models:   models,
		Metadata: metadata,
	}

	return ret, nil
}
