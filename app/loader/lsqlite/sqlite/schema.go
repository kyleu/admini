package sqlite

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema"
)

func LoadDatabaseSchema(ctx context.Context, db *database.Service, logger *zap.SugaredLogger) (*schema.Schema, error) {
	metadata := loadMetadata(db)

	tables, err := loadTables(ctx, db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "can't load tables")
	}

	err = loadColumns(ctx, tables, db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "can't load columns")
	}

	err = loadIndexes(ctx, tables, db)
	if err != nil {
		return nil, errors.Wrap(err, "can't load indexes")
	}

	err = loadForeignKeys(ctx, tables, db)
	if err != nil {
		return nil, errors.Wrap(err, "can't load foreign keys")
	}

	tables.Sort()

	ret := &schema.Schema{
		Paths:    []string{"postgres:" + db.DatabaseName},
		Models:   tables,
		Metadata: metadata,
	}

	return ret, nil
}
