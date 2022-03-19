package loader

import (
	"context"

	"admini.dev/admini/app/lib/filter"
	"admini.dev/admini/app/lib/schema"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/result"
	"admini.dev/admini/app/util"
)

type Loader interface {
	Schema(ctx context.Context) (*schema.Schema, error)

	Connection(ctx context.Context) (any, error)
	List(ctx context.Context, m *model.Model, opts *filter.Options) (*result.Result, error)
	Count(ctx context.Context, m *model.Model) (int, error)
	Get(ctx context.Context, m *model.Model, ids []any) (*result.Result, error)
	Query(ctx context.Context, enums model.Models, sql string) (*result.Result, error)
	Add(ctx context.Context, m *model.Model, changes util.ValueMap) ([]any, error)
	Save(ctx context.Context, m *model.Model, ids []any, changes util.ValueMap) ([]any, error)
	Remove(ctx context.Context, m *model.Model, fields []string, values []any, expected int) (int, error)
	Default(ctx context.Context, m *model.Model) ([]any, error)
}
