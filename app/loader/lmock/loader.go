package lmock

import (
	"github.com/kyleu/admini/app/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/kyleu/admini/app/util"

	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/result"
	"github.com/kyleu/admini/app/schema"
)

type Loader struct {
	key    string
	logger *zap.SugaredLogger
}

func NewLoader(logger *zap.SugaredLogger) func(key string, cfg []byte) (loader.Loader, error) {
	return func(key string, cfg []byte) (loader.Loader, error) {
		ret := &Loader{
			key:    key,
			logger: logger.With(zap.String("service", "loader.mock"), zap.String("source", key)),
		}
		return ret, nil
	}
}

var _ loader.Loader = (*Loader)(nil)

func (l *Loader) Connection() (interface{}, error) {
	return nil, nil
}

func (l *Loader) Schema() (*schema.Schema, error) {
	o, ok := mockData[l.key]
	if !ok {
		return nil, errors.Errorf("no mock data for [%s]", l.key)
	}

	return o.Schema, nil
}

func (l *Loader) List(m *model.Model, _ util.ParamSet) (*result.Result, error) {
	o, ok := mockData[l.key]
	if !ok {
		return nil, errors.Errorf("no mock data for [%s]", l.key)
	}

	return o.Data[m.Key], nil
}

func (l *Loader) Count(*model.Model) (int, error) {
	return 1000, nil
}

func (l *Loader) Get(*model.Model, []interface{}) (*result.Result, error) {
	return nil, errors.New("implement me")
}

func (l *Loader) Query(string) (*result.Result, error) {
	return nil, errors.New("implement me")
}

func (l *Loader) Default(*model.Model) ([]interface{}, error) {
	return nil, errors.New("implement me")
}
