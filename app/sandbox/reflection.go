package sandbox

import (
	"context"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/result"
	"go.uber.org/zap"
)

func reflectionF(ctx context.Context, st *app.State, logger *zap.SugaredLogger) (interface{}, error) {
	test := &database.PostgresParams{Host: "localhost", Port: 5432, Username: "user", Password: "pass", Database: "db", Schema: "schema", Debug: true}
	return result.FromReflection("sandbox", test)
}

var reflection = &Sandbox{Key: "reflection", Title: "Reflection", Icon: "happy", Run: reflectionF}
