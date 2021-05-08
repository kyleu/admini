package sandbox

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kyleu/admini/app/result"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/loader/lpostgres"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/queries"
)

var testbed = &Sandbox{Key: "testbed", Title: "Testbed", Run: onTestbed}

func onTestbed(st *app.State) (interface{}, error) {
	ret := map[string]interface{}{}
	sourceKey := "admini_test"
	source, err := st.Sources.Load(sourceKey)
	if err != nil {
		return nil, fmt.Errorf("can't load source: %w", err)
	}

	connInterface, err := st.Loaders.Get(schema.OriginPostgres).Connection(source.Key, source.Config)
	if err != nil {
		return nil, fmt.Errorf("can't get connection: %w", err)
	}

	conn, ok := connInterface.(*database.Service)
	if !ok {
		return nil, fmt.Errorf("invalid config object [%T]", connInterface)
	}

	run := func(key string, q string) error {
		var rows *sqlx.Rows
		rows, err = conn.Query(q, nil)
		if err != nil {
			return fmt.Errorf("can't query %v: %w", key, err)
		}
		var res *result.Result
		res, err = lpostgres.ParseResult(key, 0, q, nil, rows)
		if err != nil {
			return fmt.Errorf("can't parse result for %v: %w", key, err)
		}
		ret[key] = res
		return nil
	}

	sch := ""
	x := []struct {
		Key string
		SQL string
	}{
		{Key: "tables", SQL: queries.ListTables(sch)},
		{Key: "columns", SQL: queries.ListColumns(sch)},
		{Key: "indexes", SQL: queries.ListIndexes(sch)},
		{Key: "fks", SQL: queries.ListForeignKeys(sch)},
	}

	for _, q := range x {
		err = run(q.Key, q.SQL)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}
