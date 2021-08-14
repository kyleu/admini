// Package sandbox $PF_IGNORE$
package sandbox

import (
	"context"

	"github.com/kyleu/admini/app/loader/ldb"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/queries/qsqlite"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"
	"github.com/kyleu/admini/app/result"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/database"
)

var testbed = &Sandbox{Key: "testbed", Title: "Testbed", Icon: "code", Run: onTestbed}

func onTestbed(ctx context.Context, st *app.State, logger *zap.SugaredLogger) (interface{}, error) {
	ret := util.ValueMap{}
	sourceKey := "chinook"
	source, _ := st.Services.Sources.Load(sourceKey, true)

	if source != nil {
		load, err := st.Services.Loaders.Get(source.Type, source.Key, source.Config)
		if err != nil {
			return nil, errors.Wrap(err, "can't get loader")
		}

		connInterface, err := load.Connection(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "can't get connection")
		}

		conn, ok := connInterface.(*database.Service)
		if !ok {
			return nil, errors.Errorf("invalid config object [%T]", connInterface)
		}

		run := func(key string, q string) error {
			var rows *sqlx.Rows
			rows, err = conn.Query(ctx, q, nil)
			if err != nil {
				return errors.Wrapf(err, "can't query %s", key)
			}
			var res *result.Result
			res, err = ldb.ParseResult(key, 0, q, nil, rows, logger)
			if err != nil {
				return errors.Wrapf(err, "can't parse result for %s", key)
			}
			ret[key] = res
			return nil
		}

		sch := ""
		x := []struct {
			Key string
			SQL string
		}{
			{Key: "tables", SQL: qsqlite.ListTables(sch)},
			{Key: "columns", SQL: qsqlite.ListColumns(sch)},
			{Key: "indexes", SQL: qsqlite.ListIndexes(sch)},
			{Key: "fks", SQL: qsqlite.ListForeignKeys(sch)},
		}

		for _, q := range x {
			err = run(q.Key, q.SQL)
			if err != nil {
				return nil, err
			}
		}
	}

	return ret, nil
}
