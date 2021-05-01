package lmock

import (
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
			Models: schema.Models{
				{
					Key:  "a1x1",
					Pkg:  util.Pkg{"a", "1", "x"},
					Type: schema.ModelTypeStruct,
				},
			},
		},
		Data: map[string]*result.Result{},
	},
}
