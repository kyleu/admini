package loader

import (
	"github.com/kyleu/admini/app/result"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
)

type Loader interface {
	Schema(source string, cfg []byte) (*schema.Schema, error)

	Connection(source string, cfg []byte) (interface{}, error)
	List(source string, cfg []byte, model *schema.Model, params util.ParamSet) (*result.Result, error)
	Count(source string, cfg []byte, model *schema.Model) (int, error)
}
