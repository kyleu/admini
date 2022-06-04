// Package app $PF_IGNORE$
package app

import (
	"context"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema"
	"admini.dev/admini/app/loader"
	"admini.dev/admini/app/loader/lmock"
	"admini.dev/admini/app/loader/lmysql"
	"admini.dev/admini/app/loader/lpostgres"
	"admini.dev/admini/app/loader/lsqlite"
	"admini.dev/admini/app/project"
	"admini.dev/admini/app/source"
	"admini.dev/admini/app/util"
)

type Services struct {
	Sources  *source.Service
	Projects *project.Service
	Loaders  *loader.Service
}

func NewServices(ctx context.Context, st *State, logger util.Logger) (*Services, error) {
	ls := loader.NewService()
	ls.Set(schema.OriginMySQL, lmysql.NewLoader(ctx, logger))
	ls.Set(schema.OriginPostgres, lpostgres.NewLoader(ctx, logger))
	if database.SQLiteEnabled {
		ls.Set(schema.OriginSQLite, lsqlite.NewLoader(ctx, logger))
	}
	ls.Set(schema.OriginMock, lmock.NewLoader(ctx))
	ss := source.NewService(st.Files, ls)
	ps := project.NewService(st.Files, ss, ls)

	return &Services{Sources: ss, Projects: ps, Loaders: ls}, nil
}

func (s *Services) Close(_ context.Context, logger util.Logger) error {
	return nil
}
