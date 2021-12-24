package loader

import (
	"context"

	"github.com/kyleu/admini/app/filter"
	"github.com/kyleu/admini/app/result"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/schema/model"
	"github.com/kyleu/admini/app/util"
)

type Loader interface {
	Schema(ctx context.Context) (*schema.Schema, error)

	Connection(ctx context.Context) (interface{}, error)
	List(ctx context.Context, m *model.Model, opts *filter.Options) (*result.Result, error)
	Count(ctx context.Context, m *model.Model) (int, error)
	Get(ctx context.Context, m *model.Model, ids []interface{}) (*result.Result, error)
	Query(ctx context.Context, enums model.Models, sql string) (*result.Result, error)
	Add(ctx context.Context, m *model.Model, changes util.ValueMap) ([]interface{}, error)
	Save(ctx context.Context, m *model.Model, ids []interface{}, changes util.ValueMap) ([]interface{}, error)
	Remove(ctx context.Context, m *model.Model, fields []string, values []interface{}, expected int) (int, error)
	Default(ctx context.Context, m *model.Model) ([]interface{}, error)
}
