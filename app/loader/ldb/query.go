package ldb

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"admini.dev/app/lib/database"
	"admini.dev/app/lib/schema/model"
	"admini.dev/app/result"
)

func Query(ctx context.Context, db *database.Service, sql string, enums model.Models, logger *zap.SugaredLogger) (*result.Result, error) {
	rows, err := db.Query(ctx, sql, nil)
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
