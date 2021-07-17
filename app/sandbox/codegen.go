package sandbox

import (
	"github.com/kyleu/admini/app"
	"go.uber.org/zap"
)

var codegen = &Sandbox{Key: "codegen", Title: "Code Generation", Icon: "print", Run: func(st *app.State, logger *zap.SugaredLogger) (interface{}, error) {
	return "Work in progress...", nil
}}
