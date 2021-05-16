package loader

import (
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/result"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
)

type Loader interface {
	Schema() (*schema.Schema, error)

	Connection() (interface{}, error)
	List(m *model.Model, params util.ParamSet) (*result.Result, error)
	Count(m *model.Model) (int, error)
	Get(m *model.Model, ids []interface{}) (*result.Result, error)
	Query(sql string) (*result.Result, error)
	Default(m *model.Model) ([]interface{}, error)
}
