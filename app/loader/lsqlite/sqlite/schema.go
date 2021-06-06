package sqlite

import (
	"github.com/kyleu/admini/app/model"
	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/schema"
)

func LoadDatabaseSchema(db *database.Service, logger *zap.SugaredLogger) (*schema.Schema, error) {
	metadata := loadMetadata(db)

	scalars, err := loadScalars()
	if err != nil {
		return nil, errors.Wrap(err, "can't load scalars")
	}

	enums, err := loadEnums(db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "can't load enums")
	}

	tables, err := loadTables(enums, db, logger)
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

	models := make(model.Models, 0, len(tables)+len(enums))
	models = append(models, enums...)
	models = append(models, tables...)
	models.Sort()

	ret := &schema.Schema{
		Paths:    []string{"postgres:" + db.DatabaseName},
		Scalars:  scalars,
		Models:   models,
		Metadata: metadata,
	}

	return ret, nil
}
