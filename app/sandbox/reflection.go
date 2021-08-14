package sandbox

import (
	"context"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/result"
	"go.uber.org/zap"
)

var reflection = &Sandbox{Key: "reflection", Title: "Reflection", Icon: "happy", Run: func(ctx context.Context, st *app.State, logger *zap.SugaredLogger) (interface{}, error) {
	test := &database.PostgresParams{Host: "localhost", Port: 5432, Username: "user", Password: "pass", Database: "db", Schema: "schema", Debug: true}
	ret, err := result.FromReflection("sandbox", test)
	return ret, err
}}
