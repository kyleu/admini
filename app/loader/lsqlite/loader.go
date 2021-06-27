package lsqlite

import (
	"github.com/kyleu/admini/app/filter"
	"github.com/kyleu/admini/app/loader/ldb"
	"github.com/kyleu/admini/app/loader/lsqlite/sqlite"
	"github.com/kyleu/admini/app/result"
	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
)

type Loader struct {
	key    string
	db     *database.Service
	logger *zap.SugaredLogger
}

func NewLoader(logger *zap.SugaredLogger) func(key string, cfg []byte) (loader.Loader, error) {
	return func(key string, cfg []byte) (loader.Loader, error) {
		log := logger.With(zap.String("service", "loader.sqlite"), zap.String("source", key))
		db, err := openDatabase(cfg, log)
		if err != nil {
			return nil, errors.Wrap(err, "error opening database")
		}
		return &Loader{key: key, db: db, logger: log}, nil
	}
}

var _ loader.Loader = (*Loader)(nil)

func (l *Loader) Schema() (*schema.Schema, error) {
	return sqlite.LoadDatabaseSchema(l.db, l.logger)
}

func (l *Loader) Connection() (interface{}, error) {
	return l.db, nil
}

func (l *Loader) List(m *model.Model, opts *filter.Options) (*result.Result, error) {
	return ldb.List(l.db, m, opts, l.logger)
}

func (l *Loader) Count(m *model.Model) (int, error) {
	return ldb.Count(l.db, m)
}

func (l *Loader) Get(m *model.Model, ids []interface{}) (*result.Result, error) {
	return ldb.Get(l.db, m, ids, l.logger)
}

func (l *Loader) Query(sql string) (*result.Result, error) {
	return ldb.Query(l.db, sql, l.logger)
}

func (l *Loader) Add(m *model.Model, changes util.ValueMap) ([]interface{}, error) {
	return ldb.Add(l.db, m, changes, l.logger)
}

func (l *Loader) Save(m *model.Model, ids []interface{}, changes util.ValueMap) ([]interface{}, error) {
	return ldb.Save(l.db, m, ids, changes, l.logger)
}

func (l *Loader) Remove(m *model.Model, fields []string, values []interface{}, expected int) (int, error) {
	return ldb.Remove(l.db, m, fields, values, expected, l.logger)
}

func (l *Loader) Default(m *model.Model) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(m.Fields))
	for _, f := range m.Fields {
		ret = append(ret, f.Default)
	}
	return ret, nil
}

func LoadConfig(cfg []byte) (*database.SQLiteParams, error) {
	params := &database.SQLiteParams{}
	err := util.FromJSON(cfg, params)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing database config")
	}
	return params, nil
}

func openDatabase(cfg []byte, logger *zap.SugaredLogger) (*database.Service, error) {
	params, err := LoadConfig(cfg)
	if err != nil {
		return nil, err
	}
	db, err := database.OpenSQLiteDatabase(params, logger)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database")
	}

	_, err = db.SingleInt("select 1 as x", nil)
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to database")
	}
	return db, nil
}
