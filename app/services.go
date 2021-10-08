// Package app $PF_IGNORE$
package app

import (
	"context"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/loader/lmock"
	"github.com/kyleu/admini/app/loader/lmysql"
	"github.com/kyleu/admini/app/loader/lpostgres"
	"github.com/kyleu/admini/app/loader/lsqlite"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/source"
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
