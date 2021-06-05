package lmock

import (
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/result"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
)

type mockObj struct {
	Schema *schema.Schema
	Data   map[string]*result.Result
}

var mockData = map[string]*mockObj{
	"mock": {
		Schema: &schema.Schema{
			Models: model.Models{
				{
					Key:  "a1x1",
					Pkg:  util.Pkg{"a", "1", "x"},
					Type: model.TypeStruct,
				},
			},
		},
		Data: map[string]*result.Result{
			"a1x1": result.NewResult("mock", 0, "mock", nil, nil, nil),
		},
	},
}
