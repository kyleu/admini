package lpostgres

import (
	"strings"

	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/source/postgres"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
)

type Loader struct {
	key    string
	db     *database.Service
	logger *zap.SugaredLogger
}

func NewLoader(logger *zap.SugaredLogger, debug bool) func(key string, cfg []byte) (loader.Loader, error) {
	return func(key string, cfg []byte) (loader.Loader, error) {
		log := logger.With(zap.String("service", "loader.postgres"), zap.String("source", key))
		db, err := openDatabase(cfg, debug, log)
		if err != nil {
			return nil, errors.Wrap(err, "error opening database")
		}
		return &Loader{key: key, db: db, logger: log}, nil
	}
}

var _ loader.Loader = (*Loader)(nil)

func (l *Loader) Connection() (interface{}, error) {
	return l.db, nil
}

func (l *Loader) Schema() (*schema.Schema, error) {
	return postgres.LoadDatabaseSchema(l.db, l.logger)
}

func (l *Loader) Default(m *model.Model) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(m.Fields))
	for _, f := range m.Fields {
		ret = append(ret, f.Default)
	}
	return ret, nil
}

func openDatabase(cfg []byte, debug bool, logger *zap.SugaredLogger) (*database.Service, error) {
	config := &database.DBParams{}
	err := util.FromJSON(cfg, config)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing database config")
	}
	config.Debug = debug
	db, err := database.OpenDatabase(config, logger)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database")
	}

	_, err = db.SingleInt("select 1 as x", nil)
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to database")
	}
	return db, nil
}

func forTable(m *model.Model) (string, string) {
	cols := make([]string, 0, len(m.Fields))
	for _, f := range m.Fields {
		cols = append(cols, "\""+f.Key+"\"")
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
