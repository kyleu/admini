package lpostgres

import (
	"database/sql"

	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/field"

	"github.com/jmoiron/sqlx"
	"github.com/kyleu/admini/app/result"
	"github.com/kyleu/admini/app/source/postgres"
)

func ParseResult(title string, count int, q string, timing *result.Timing, rows *sqlx.Rows, logger *zap.SugaredLogger) (*result.Result, error) {
	fields, err := parseColumns(rows, logger)
	if err != nil {
		return nil, errors.Wrap(err, "error processing columns")
	}
	return ParseResultFields(title, count, q, timing, fields, rows)
}

func ParseResultFields(title string, count int, q string, timing *result.Timing, fields field.Fields, rows *sqlx.Rows) (*result.Result, error) {
	data, err := parseRows(rows)
	if err != nil {
		return nil, errors.Wrap(err, "error processing database rows")
	}
	if count == 0 {
		count = len(data)
	}
	ret := result.NewResult(title, count, q, fields, data, timing)
	return ret, nil
}

func parseColumns(rows *sqlx.Rows, logger *zap.SugaredLogger) (field.Fields, error) {
	cts, err := rows.ColumnTypes()
	if err != nil {
		return nil, errors.Wrap(err, "unable to determine column types")
	}

	fs := make(field.Fields, 0, len(cts))
	for _, ct := range cts {
		f, err := fieldFor(ct, logger)
		if err != nil {
			return nil, errors.Wrap(err, "unable to parse column type")
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
			return nil, errors.Wrap(err, "unable to scan results")
		}

		data = append(data, d)
	}
	return data, nil
}

func fieldFor(ct *sql.ColumnType, logger *zap.SugaredLogger) (*field.Field, error) {
	return &field.Field{
		Key:  ct.Name(),
		Type: postgres.TypeForName(ct.DatabaseTypeName(), logger),
	}, nil
}
