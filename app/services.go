// Package app $PF_IGNORE$
package app

import (
	"context"

	"admini.dev/app/lib/database"
	"admini.dev/app/lib/schema"
	"admini.dev/app/loader"
	"admini.dev/app/loader/lmock"
	"admini.dev/app/loader/lmysql"
	"admini.dev/app/loader/lpostgres"
	"admini.dev/app/loader/lsqlite"
	"admini.dev/app/project"
	"admini.dev/app/source"
)

type Services struct {
	Sources  *source.Service
	Projects *project.Service
	Loaders  *loader.Service
}

func NewServices(ctx context.Context, st *State) (*Services, error) {
	ls := loader.NewService()
	ls.Set(schema.OriginMySQL, lmysql.NewLoader(ctx, st.Logger))
	ls.Set(schema.OriginPostgres, lpostgres.NewLoader(ctx, st.Logger))
	if database.SQLiteEnabled {
		ls.Set(schema.OriginSQLite, lsqlite.NewLoader(ctx, st.Logger))
	}
	ls.Set(schema.OriginMock, lmock.NewLoader(ctx, st.Logger))
	ss := source.NewService(st.Files, ls, st.Logger)
	ps := project.NewService(st.Files, ss, ls, st.Logger)

	return &Services{Sources: ss, Projects: ps, Loaders: ls}, nil
}
