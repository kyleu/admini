package postgres

import (
	"encoding/json"
	"fmt"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
)

type Loader struct{}

func (l *Loader) GetSchema(configJSON json.RawMessage) (*schema.Schema, error) {
	config := &database.DBParams{}
	err := util.FromJSON(configJSON, config)
	if err != nil {
		return nil, fmt.Errorf("error parsing database config: %w", err)
	}

	db, err := database.OpenDatabase(config)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	err = testConnect(db)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return loadSchema(config, db)
}

func testConnect(db *database.Service) error {
	return db.Get(&database.Count{}, "select 1 as c", nil)
}

func loadSchema(params *database.DBParams, db *database.Service) (*schema.Schema, error) {
	var errs []string
	addErr := func(err error) {
		errs = append(errs, fmt.Sprintf("%+v", err))
	}

	metadata, err := loadMetadata(params.Database, db)
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
		Paths:    []string{"postgres:" + params.Database},
		Scalars:  scalars,
		Models:   models,
		Errors:   errs,
		Metadata: metadata,
	}

	return ret, nil
}
