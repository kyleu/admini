package postgres

import (
	"fmt"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/schema"
)

func LoadDatabaseSchema(db *database.Service) (*schema.Schema, error) {
	var errs []string
	addErr := func(err error) {
		errs = append(errs, fmt.Sprintf("%+v", err))
	}

	metadata, err := loadMetadata(db)
	if err != nil {
		addErr(fmt.Errorf("can't load metadata: %w", err))
	}

	scalars, err := loadScalars(db)
	if err != nil {
		addErr(fmt.Errorf("can't load scalars: %w", err))
	}

	models, err := loadTables(db)
	if err != nil {
		addErr(fmt.Errorf("can't load tables: %w", err))
	}

	err = loadColumns(models, db)
	if err != nil {
		addErr(fmt.Errorf("can't load columns: %w", err))
	}

	err = loadIndexes(models, db)
	if err != nil {
		addErr(fmt.Errorf("can't load indexes: %w", err))
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
