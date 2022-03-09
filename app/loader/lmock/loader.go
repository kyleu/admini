package lmock

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"admini.dev/admini/app/lib/filter"
	"admini.dev/admini/app/lib/schema"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/loader"
	"admini.dev/admini/app/result"
	"admini.dev/admini/app/util"
)

type Loader struct {
	key    string
	logger *zap.SugaredLogger
}

func NewLoader(_ context.Context, logger *zap.SugaredLogger) func(key string, cfg []byte) (loader.Loader, error) {
	return func(key string, cfg []byte) (loader.Loader, error) {
		ret := &Loader{
			key:    key,
			logger: logger.With(zap.String("service", "loader.mock"), zap.String("source", key)),
		}
		return ret, nil
	}
}

var _ loader.Loader = (*Loader)(nil)

func (l *Loader) Schema(context.Context) (*schema.Schema, error) {
	o, ok := mockData[l.key]
	if !ok {
		return nil, errors.Errorf("no mock data for [%s]", l.key)
	}

	return o.Schema, nil
}

func (l *Loader) Connection(context.Context) (interface{}, error) {
	return nil, nil
}

func (l *Loader) List(_ context.Context, m *model.Model, _ *filter.Options) (*result.Result, error) {
	o, ok := mockData[l.key]
	if !ok {
		return nil, errors.Errorf("no mock data for [%s]", l.key)
	}

	return o.Data[m.Key], nil
}

func (l *Loader) Count(context.Context, *model.Model) (int, error) {
	return 1000, nil
}

func (l *Loader) Get(context.Context, *model.Model, []interface{}) (*result.Result, error) {
	return nil, errors.New("mock not implemented")
}

func (l *Loader) Query(context.Context, model.Models, string) (*result.Result, error) {
	return nil, errors.New("mock not implemented")
}

func (l *Loader) Add(context.Context, *model.Model, util.ValueMap) ([]interface{}, error) {
	return nil, errors.New("mock not implemented")
}

func (l *Loader) Save(context.Context, *model.Model, []interface{}, util.ValueMap) ([]interface{}, error) {
	return nil, errors.New("mock not implemented")
}

func (l *Loader) Remove(context.Context, *model.Model, []string, []interface{}, int) (int, error) {
	return 0, errors.New("mock not implemented")
}

func (l *Loader) Default(context.Context, *model.Model) ([]interface{}, error) {
	return nil, errors.New("mock not implemented")
}
