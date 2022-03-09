package lmock

import (
	"admini.dev/admini/app/lib/schema"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/result"
	"admini.dev/admini/app/util"
)

type mockObj struct {
	Schema *schema.Schema
	Data   map[string]*result.Result
}

var mockData = map[string]*mockObj{
	"mock": {
		Schema: &schema.Schema{
			Models: model.Models{
				&model.Model{Key: "a1x1", Pkg: util.Pkg{"a", "1", "x"}, Type: model.TypeStruct},
			},
		},
		Data: map[string]*result.Result{
			"a1x1": result.NewResult("mock", 0, "mock", nil, nil, nil),
		},
	},
}
