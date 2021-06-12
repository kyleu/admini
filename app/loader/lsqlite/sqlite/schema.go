package sqlite

import (
	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/schema"
)

func LoadDatabaseSchema(db *database.Service, logger *zap.SugaredLogger) (*schema.Schema, error) {
	metadata := loadMetadata(db)

	tables, err := loadTables(db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "can't load tables")
	}

	err = loadColumns(tables, db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "can't load columns")
	}

	err = loadIndexes(tables, db)
	if err != nil {
		return nil, errors.Wrap(err, "can't load indexes")
	}

	err = loadForeignKeys(tables, db)
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
