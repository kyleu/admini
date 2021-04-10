package sandbox

import (
	"github.com/kyleu/admini/app/schema"
)

type Sandbox struct {
	Key string
	Run func() (interface{}, error)
}

var testbed = &Sandbox{Key: "testbed", Run: func() (interface{}, error) {
	return "OK", nil
}}

var codegen = &Sandbox{Key: "codegen", Run: func() (interface{}, error) {
	return schema.GetExample(), nil
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

var AllSandboxes = Sandboxes{codegen, testbed}
