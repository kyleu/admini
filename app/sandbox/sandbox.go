package sandbox

import (
	"github.com/kyleu/admini/app"
)

type Sandbox struct {
	Key   string                                   `json:"key,omitempty"`
	Title string                                   `json:"title,omitempty"`
	Run   func(st *app.State) (interface{}, error) `json:"-"`
}

var codegen = &Sandbox{Key: "codegen", Title: "Code Generation", Run: func(st *app.State) (interface{}, error) {
	return "TODO", nil
}}

var sources = &Sandbox{Key: "sources", Title: "Data Sources", Run: func(st *app.State) (interface{}, error) {
	return st.Sources.List()
}}

var testbed = &Sandbox{Key: "testbed", Title: "Testbed", Run: func(st *app.State) (interface{}, error) {
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
