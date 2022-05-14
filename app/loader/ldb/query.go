package ldb

import (
	"context"

	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/result"
	"admini.dev/admini/app/util"
)

func Query(ctx context.Context, db *database.Service, sql string, enums model.Models, logger util.Logger) (*result.Result, error) {
	rows, err := db.Query(ctx, sql, nil, logger)
	if err != nil {
		return nil, errors.Wrap(err, "error running SQL")
	}
	defer func() { _ = rows.Close() }()

	var timing *result.Timing
	ret, err := ParseResult("SQL result", 0, sql, timing, rows, enums, logger)
	if err != nil {
		return nil, errors.Wrap(err, "error constructing result for SQL")
	}

	return ret, nil
}
