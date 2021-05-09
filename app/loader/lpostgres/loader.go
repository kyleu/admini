package lpostgres

import (
	"github.com/pkg/errors"
	"strings"

	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/model"
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
		return nil, errors.Wrap(err, "error opening database")
	}

	return db, nil
}

func (l *Loader) Schema(source string, cfg []byte) (*schema.Schema, error) {
	db, err := l.openDatabase(source, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database")
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
		return nil, errors.Wrap(err, "error parsing database config")
	}
	db, err := database.OpenDatabase(config)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database")
	}

	_, err = db.SingleInt("select 1 as x", nil)
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to database")
	}

	l.cache[source] = db
	return db, nil
}

func forTable(m *model.Model) (string, string) {
	cols := make([]string, 0, len(m.Fields))
	for _, f := range m.Fields {
		def := "\"" + f.Key + "\""
		cols = append(cols, def)
	}
	tbl := "\"" + m.Key + "\""
	if len(m.Pkg) > 0 {
		l := m.Pkg.Last()
		if l != publicSchema {
			tbl = "\"" + l + "\"." + tbl
		}
	}
	return strings.Join(cols, ", "), tbl
}
