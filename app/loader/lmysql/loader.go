package lmysql

import (
	"context"

	"github.com/kyleu/admini/app/filter"
	"github.com/kyleu/admini/app/loader/ldb"
	"github.com/kyleu/admini/app/loader/lmysql/mysql"
	"github.com/kyleu/admini/app/result"
	"github.com/kyleu/admini/app/schema/model"
	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
)

type Loader struct {
	key    string
	db     *database.Service
	logger *zap.SugaredLogger
}

func NewLoader(ctx context.Context, logger *zap.SugaredLogger) func(key string, cfg []byte) (loader.Loader, error) {
	return func(key string, cfg []byte) (loader.Loader, error) {
		log := logger.With(zap.String("service", "loader.mysql"), zap.String("source", key))
		db, err := openDatabase(ctx, key, cfg, log)
		if err != nil {
			return nil, errors.Wrap(err, "error opening database")
		}
		return &Loader{key: key, db: db, logger: log}, nil
	}
}

var _ loader.Loader = (*Loader)(nil)

func (l *Loader) Schema(ctx context.Context) (*schema.Schema, error) {
	return mysql.LoadDatabaseSchema(ctx, l.db, l.logger)
}

func (l *Loader) Connection(ctx context.Context) (interface{}, error) {
	return l.db, nil
}

func (l *Loader) List(ctx context.Context, m *model.Model, opts *filter.Options) (*result.Result, error) {
	return ldb.List(ctx, l.db, m, opts)
}

func (l *Loader) Count(ctx context.Context, m *model.Model) (int, error) {
	return ldb.Count(ctx, l.db, m)
}

func (l *Loader) Get(ctx context.Context, m *model.Model, ids []interface{}) (*result.Result, error) {
	return ldb.Get(ctx, l.db, m, ids, l.logger)
}

func (l *Loader) Query(ctx context.Context, enums model.Models, sql string) (*result.Result, error) {
	return ldb.Query(ctx, l.db, sql, enums, l.logger)
}

func (l *Loader) Add(ctx context.Context, m *model.Model, changes util.ValueMap) ([]interface{}, error) {
	return ldb.Add(ctx, l.db, m, changes, l.logger)
}

func (l *Loader) Save(ctx context.Context, m *model.Model, ids []interface{}, changes util.ValueMap) ([]interface{}, error) {
	return ldb.Save(ctx, l.db, m, ids, changes, l.logger)
}

func (l *Loader) Remove(ctx context.Context, m *model.Model, fields []string, values []interface{}, expected int) (int, error) {
	return ldb.Remove(ctx, l.db, m, fields, values, expected, l.logger)
}

func (l *Loader) Default(ctx context.Context, m *model.Model) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(m.Fields))
	for _, f := range m.Fields {
		ret = append(ret, f.DefaultClean())
	}
	return ret, nil
}

func LoadConfig(cfg []byte) (*database.MySQLParams, error) {
	params := &database.MySQLParams{}
	err := util.FromJSON(cfg, params)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing database config")
	}
	return params, nil
}

func openDatabase(ctx context.Context, key string, cfg []byte, logger *zap.SugaredLogger) (*database.Service, error) {
	params, err := LoadConfig(cfg)
	if err != nil {
		return nil, err
	}
	db, err := database.OpenMySQLDatabase(ctx, key, params, logger)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database")
	}

	_, err = db.SingleInt(ctx, "select 1 as x", nil)
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to database")
	}
	return db, nil
}
