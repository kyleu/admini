package lpostgres

import (
	"fmt"

	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/source/postgres"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
)

type Loader struct {
	cache map[string]*database.Service
}

func NewLoader() *Loader {
	return &Loader{cache: map[string]*database.Service{}}
}

var _ loader.Loader = (*Loader)(nil)

func (l *Loader) Connection(source string, cfg []byte) (interface{}, error) {
	db, err := l.openDatabase(source, cfg)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	return db, nil
}

func (l *Loader) Schema(source string, cfg []byte) (*schema.Schema, error) {
	db, err := l.openDatabase(source, cfg)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	return postgres.LoadDatabaseSchema(db)
}

func (l *Loader) openDatabase(source string, cfg []byte) (*database.Service, error) {
	x, ok := l.cache[source]
	if ok {
		return x, nil
	}
	config := &database.DBParams{}
	err := util.FromJSON(cfg, config)
	if err != nil {
		return nil, fmt.Errorf("error parsing database config: %w", err)
	}
	db, err := database.OpenDatabase(config)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	_, err = db.SingleInt("select 1 as x", nil)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	l.cache[source] = db
	return db, nil
}
