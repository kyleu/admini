package sqlite

import (
	"context"

	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema"
	"admini.dev/admini/app/util"
)

func LoadDatabaseSchema(ctx context.Context, db *database.Service, logger util.Logger) (*schema.Schema, error) {
	metadata := loadMetadata(db)

	tables, err := loadTables(ctx, db, logger)
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

	tables.Sort()

	ret := &schema.Schema{
		Paths:    []string{"postgres:" + db.DatabaseName},
		Models:   tables,
		Metadata: metadata,
	}

	return ret, nil
}
