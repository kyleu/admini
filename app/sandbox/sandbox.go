package sandbox

import (
	"github.com/kyleu/admini/app/schema"
)

type Sandbox struct {
	Key string
	Title string
	Run func() (interface{}, error)
}

var codegen = &Sandbox{Key: "codegen", Title: "Code Generation", Run: func() (interface{}, error) {
	return schema.GetExample(), nil
}}

var datasources = &Sandbox{Key: "datasources", Title: "Data Sources", Run: func() (interface{}, error) {
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

var AllSandboxes = Sandboxes{codegen, datasources, testbed}
