package sandbox

import (
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
)

type Sandbox struct {
	Key   string
	Title string
	Run   func() (interface{}, error)
}

var codegen = &Sandbox{Key: "codegen", Title: "Code Generation", Run: func() (interface{}, error) {
	var sch schema.Fields
	err := util.FromJSON(util.ToJSONBytes(schema.GetExample(), true), &sch)
	if err != nil {
		return nil, err
	}
	return sch, nil
}}

var sources = &Sandbox{Key: "sources", Title: "Data Sources", Run: func() (interface{}, error) {
	return schema.GetExample(), nil
}}

var testbed = &Sandbox{Key: "testbed", Title: "Testbed", Run: func() (interface{}, error) {
	return "OK", nil
}}

type Sandboxes []*Sandbox

func (s Sandboxes) Get(key string) *Sandbox {
	for _, v := range s {
		if v.Key == key {
			return v
		}
	}
	return nil
}

var AllSandboxes = Sandboxes{codegen, sources, testbed}
