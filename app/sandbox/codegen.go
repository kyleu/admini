package sandbox

import (
	"context"

	"github.com/kyleu/admini/app"
	"go.uber.org/zap"
)

func codegenF(ctx context.Context, st *app.State, logger *zap.SugaredLogger) (interface{}, error) {
	return "Work in progress...", nil
}

var codegen = &Sandbox{Key: "codegen", Title: "Code Generation", Icon: "print", Run: codegenF}
