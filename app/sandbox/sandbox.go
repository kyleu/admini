package sandbox

import (
	"github.com/kyleu/admini/app"
	"go.uber.org/zap"
)

type Sandbox struct {
	Key   string                                                              `json:"key,omitempty"`
	Title string                                                              `json:"title,omitempty"`
	Run   func(st *app.State, logger *zap.SugaredLogger) (interface{}, error) `json:"-"`
}

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

var codegen = &Sandbox{Key: "codegen", Title: "Code Generation", Run: func(st *app.State, logger *zap.SugaredLogger) (interface{}, error) {
	return "TODO", nil
}}

var sources = &Sandbox{Key: "sources", Title: "Data Sources", Run: func(st *app.State, logger *zap.SugaredLogger) (interface{}, error) {
	return st.Sources.List()
}}
