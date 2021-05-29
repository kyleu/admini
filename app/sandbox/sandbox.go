package sandbox

import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/result"
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

var AllSandboxes = Sandboxes{codegen, reflection, testbed}

var codegen = &Sandbox{Key: "codegen", Title: "Code Generation", Run: func(st *app.State, logger *zap.SugaredLogger) (interface{}, error) {
	return "TODO", nil
}}

var reflection = &Sandbox{Key: "reflection", Title: "Reflection", Run: func(st *app.State, logger *zap.SugaredLogger) (interface{}, error) {
	test := &database.DBParams{Host: "localhost", Port: 5432, Username: "user", Password: "pass", Database: "db", Schema: "schema", Debug: true}
	ret, err := result.FromReflection("sandbox", test)
	return ret, err
}}
