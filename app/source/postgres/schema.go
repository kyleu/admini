package postgres

import (
	"github.com/kyleu/admini/app/model"
	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/schema"
)

func LoadDatabaseSchema(db *database.Service, logger *zap.SugaredLogger) (*schema.Schema, error) {
	metadata, err := loadMetadata(db)
	if err != nil {
		return nil, errors.Wrap(err, "can't load metadata")
	}

	scalars, err := loadScalars(db)
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
	for _, e := range enums {
		models = append(models, e)
	}
	for _, t := range tables {
		models = append(models, t)
	}
	models.Sort()

	ret := &schema.Schema{
		Paths:    []string{"postgres:" + db.DatabaseName},
		Scalars:  scalars,
		Models:   models,
		Metadata: metadata,
	}

	return ret, nil
}
