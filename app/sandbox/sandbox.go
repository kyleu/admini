package sandbox

import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/result"
	"go.uber.org/zap"
)

type runFn func(st *app.State, logger *zap.SugaredLogger) (interface{}, error)

type Sandbox struct {
	Key   string `json:"key,omitempty"`
	Title string `json:"title,omitempty"`
	Icon  string `json:"icon,omitempty"`
	Run   runFn  `json:"-"`
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

var codegen = &Sandbox{Key: "codegen", Title: "Code Generation", Icon: "print", Run: func(st *app.State, logger *zap.SugaredLogger) (interface{}, error) {
	return "Work in progress...", nil
}}

var reflection = &Sandbox{Key: "reflection", Title: "Reflection", Icon: "happy", Run: func(st *app.State, logger *zap.SugaredLogger) (interface{}, error) {
	test := &database.PostgresParams{Host: "localhost", Port: 5432, Username: "user", Password: "pass", Database: "db", Schema: "schema", Debug: true}
	ret, err := result.FromReflection("sandbox", test)
	return ret, err
}}
