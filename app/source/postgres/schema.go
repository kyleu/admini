package postgres

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/schema"
)

func LoadDatabaseSchema(db *database.Service, logger *zap.SugaredLogger) (*schema.Schema, error) {
	var errs []string
	addErr := func(err error) {
		errs = append(errs, fmt.Sprintf("%+v", err))
	}

	metadata, err := loadMetadata(db)
	if err != nil {
		addErr(errors.Wrap(err, "can't load metadata"))
	}

	scalars, err := loadScalars(db)
	if err != nil {
		addErr(errors.Wrap(err, "can't load scalars"))
	}

	models, err := loadTables(db, logger)
	if err != nil {
		addErr(errors.Wrap(err, "can't load tables"))
	}

	err = loadColumns(models, db, logger)
	if err != nil {
		addErr(errors.Wrap(err, "can't load columns"))
	}

	err = loadIndexes(models, db)
	if err != nil {
		addErr(errors.Wrap(err, "can't load indexes"))
	}

	err = loadForeignKeys(models, db)
	if err != nil {
		addErr(errors.Wrap(err, "can't load foreign keys"))
	}

	ret := &schema.Schema{
		Paths:    []string{"postgres:" + db.DatabaseName},
		Scalars:  scalars,
		Models:   models,
		Errors:   errs,
		Metadata: metadata,
	}

	return ret, nil
}
