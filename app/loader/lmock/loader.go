package lmock

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/model"

	"github.com/kyleu/admini/app/util"

	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/result"
	"github.com/kyleu/admini/app/schema"
)

type Loader struct{}

func NewLoader() *Loader {
	return &Loader{}
}

var _ loader.Loader = (*Loader)(nil)

func (l *Loader) Connection(source string, cfg []byte) (interface{}, error) {
	return nil, nil
}

func (l *Loader) Schema(source string, cfg []byte) (*schema.Schema, error) {
	o, ok := mockData[source]
	if !ok {
		return nil, errors.New(fmt.Sprintf("no mock data for [%v]", source))
	}

	return o.Schema, nil
}

func (l *Loader) List(source string, cfg []byte, m *model.Model, params util.ParamSet) (*result.Result, error) {
	o, ok := mockData[source]
	if !ok {
		return nil, errors.New(fmt.Sprintf("no mock data for [%v]", source))
	}

	return o.Data[m.Key], nil
}

func (l *Loader) Count(source string, cfg []byte, m *model.Model) (int, error) {
	return 1000, nil
}

func (l *Loader) Get(key string, config []byte, m *model.Model, ids []interface{}) (*result.Result, error) {
	panic("implement me")
}
