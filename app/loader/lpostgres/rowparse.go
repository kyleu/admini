package lpostgres

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kyleu/admini/app/result"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/source/postgres"
)

func NewResult(title string, q string, timing *result.Timing, rows *sqlx.Rows) (*result.Result, error) {
	fields, data, err := parseRows(rows)
	if err != nil {
		return nil, fmt.Errorf("error processing database rows: %w", err)
	}

	ret := &result.Result{Title: title, Query: q, Fields: fields, Data: data, Timing: timing}
	return ret, nil
}

func parseRows(rows *sqlx.Rows) (schema.Fields, [][]interface{}, error) {
	cts, err := rows.ColumnTypes()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to determine column types: %w", err)
	}

	fs := make(schema.Fields, 0, len(cts))
	for _, ct := range cts {
		f, err := fieldFor(ct)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse column type: %w", err)
		}
		fs = append(fs, f)
	}

	data := [][]interface{}{}

	for rows.Next() {
		d, err := rows.SliceScan()
		if err != nil {
			return nil, nil, fmt.Errorf("unable to scan results: %w", err)
		}

		data = append(data, d)
	}

	return fs, data, nil
}

func fieldFor(ct *sql.ColumnType) (*schema.Field, error) {
	return &schema.Field{
		Key:     ct.Name(),
		Type:    postgres.TypeForName(ct.DatabaseTypeName()),
	}, nil
}
