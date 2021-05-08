package lpostgres

import (
	"database/sql"
	"fmt"

	"github.com/kyleu/admini/app/field"

	"github.com/jmoiron/sqlx"
	"github.com/kyleu/admini/app/result"
	"github.com/kyleu/admini/app/source/postgres"
)

func ParseResult(title string, count int, q string, timing *result.Timing, rows *sqlx.Rows) (*result.Result, error) {
	fields, err := parseColumns(rows)
	if err != nil {
		return nil, fmt.Errorf("error processing columns: %w", err)
	}
	return ParseResultFields(title, count, q, timing, fields, rows)
}

func ParseResultFields(title string, count int, q string, timing *result.Timing, fields field.Fields, rows *sqlx.Rows) (*result.Result, error) {
	data, err := parseRows(rows)
	if err != nil {
		return nil, fmt.Errorf("error processing database rows: %w", err)
	}
	if count == 0 {
		count = len(data)
	}
	ret := result.NewResult(title, count, q, fields, data, timing)
	return ret, nil
}

func parseColumns(rows *sqlx.Rows) (field.Fields, error) {
	cts, err := rows.ColumnTypes()
	if err != nil {
		return nil, fmt.Errorf("unable to determine column types: %w", err)
	}

	fs := make(field.Fields, 0, len(cts))
	for _, ct := range cts {
		f, err := fieldFor(ct)
		if err != nil {
			return nil, fmt.Errorf("unable to parse column type: %w", err)
		}
		fs = append(fs, f)
	}
	return fs, nil
}

func parseRows(rows *sqlx.Rows) ([][]interface{}, error) {
	data := [][]interface{}{}
	for rows.Next() {
		d, err := rows.SliceScan()
		if err != nil {
			return nil, fmt.Errorf("unable to scan results: %w", err)
		}

		data = append(data, d)
	}
	return data, nil
}

func fieldFor(ct *sql.ColumnType) (*field.Field, error) {
	return &field.Field{
		Key:  ct.Name(),
		Type: postgres.TypeForName(ct.DatabaseTypeName()),
	}, nil
}
