package lmock

import (
	"fmt"

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
		return nil, fmt.Errorf("no mock data for [%v]", source)
	}

	return o.Schema, nil
}

func (l *Loader) List(source string, cfg []byte, model *schema.Model, params util.ParamSet) (*result.Result, error) {
	o, ok := mockData[source]
	if !ok {
		return nil, fmt.Errorf("no mock data for [%v]", source)
	}

	return o.Data[model.Key], nil
}

func (l *Loader) Count(source string, cfg []byte, model *schema.Model) (int, error) {
	return 1000, nil
}
