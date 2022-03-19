package lpostgres

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/filter"
	"admini.dev/admini/app/lib/schema"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/loader"
	"admini.dev/admini/app/loader/ldb"
	"admini.dev/admini/app/loader/lpostgres/postgres"
	"admini.dev/admini/app/result"
	"admini.dev/admini/app/util"
)

type Loader struct {
	key    string
	db     *database.Service
	logger *zap.SugaredLogger
}

func NewLoader(ctx context.Context, logger *zap.SugaredLogger) func(key string, cfg []byte) (loader.Loader, error) {
	return func(key string, cfg []byte) (loader.Loader, error) {
		log := logger.With(zap.String("service", "loader.postgres"), zap.String("source", key))
		db, err := openDatabase(ctx, key, cfg, log)
		if err != nil {
			return nil, errors.Wrap(err, "error opening database")
		}
		return &Loader{key: key, db: db, logger: log}, nil
	}
}

var _ loader.Loader = (*Loader)(nil)

func (l *Loader) Schema(ctx context.Context) (*schema.Schema, error) {
	return postgres.LoadDatabaseSchema(ctx, l.db, l.logger)
}

func (l *Loader) Connection(ctx context.Context) (any, error) {
	return l.db, nil
}

func (l *Loader) List(ctx context.Context, m *model.Model, opts *filter.Options) (*result.Result, error) {
	return ldb.List(ctx, l.db, m, opts)
}

func (l *Loader) Count(ctx context.Context, m *model.Model) (int, error) {
	return ldb.Count(ctx, l.db, m)
}

func (l *Loader) Get(ctx context.Context, m *model.Model, ids []any) (*result.Result, error) {
	return ldb.Get(ctx, l.db, m, ids, l.logger)
}

func (l *Loader) Query(ctx context.Context, enums model.Models, sql string) (*result.Result, error) {
	return ldb.Query(ctx, l.db, sql, enums, l.logger)
}

func (l *Loader) Add(ctx context.Context, m *model.Model, changes util.ValueMap) ([]any, error) {
	return ldb.Add(ctx, l.db, m, changes, l.logger)
}

func (l *Loader) Save(ctx context.Context, m *model.Model, ids []any, changes util.ValueMap) ([]any, error) {
	return ldb.Save(ctx, l.db, m, ids, changes, l.logger)
}

func (l *Loader) Remove(ctx context.Context, m *model.Model, fields []string, values []any, expected int) (int, error) {
	return ldb.Remove(ctx, l.db, m, fields, values, expected, l.logger)
}

func (l *Loader) Default(ctx context.Context, m *model.Model) ([]any, error) {
	ret := make([]any, 0, len(m.Fields))
	for _, f := range m.Fields {
		ret = append(ret, f.DefaultClean())
	}
	return ret, nil
}

func LoadConfig(cfg []byte) (*database.PostgresParams, error) {
	params := &database.PostgresParams{}
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
	db, err := database.OpenPostgresDatabase(ctx, key, params, logger)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database")
	}

	_, err = db.SingleInt(ctx, "select 1 as x", nil)
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to database")
	}
	return db, nil
}
