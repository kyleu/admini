package sandbox

import (
	"github.com/kyleu/admini/app/loader/ldb"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"
	"github.com/kyleu/admini/app/result"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/queries/qpostgres"
)

var testbed = &Sandbox{Key: "testbed", Title: "Testbed", Icon: "social", Run: onTestbed}

func onTestbed(st *app.State, logger *zap.SugaredLogger) (interface{}, error) {
	ret := util.ValueMap{}
	sourceKey := "rituals.dev"
	source, _ := st.Sources.Load(sourceKey, true)

	if source != nil {
		load, err := st.Loaders.Get(source.Type, source.Key, source.Config)
		if err != nil {
			return nil, errors.Wrap(err, "can't get loader")
		}

		connInterface, err := load.Connection()
		if err != nil {
			return nil, errors.Wrap(err, "can't get connection")
		}

		conn, ok := connInterface.(*database.Service)
		if !ok {
			return nil, errors.Errorf("invalid config object [%T]", connInterface)
		}

		run := func(key string, q string) error {
			var rows *sqlx.Rows
			rows, err = conn.Query(q, nil)
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
			{Key: "types", SQL: qpostgres.ListTypes(sch)},
			{Key: "tables", SQL: qpostgres.ListTables(sch)},
			{Key: "columns", SQL: qpostgres.ListColumns(sch)},
			{Key: "indexes", SQL: qpostgres.ListIndexes(sch)},
			{Key: "fks", SQL: qpostgres.ListForeignKeys(sch)},
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
