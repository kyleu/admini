package ldb

import (
	"context"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/result"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Query(ctx context.Context, db *database.Service, sql string, logger *zap.SugaredLogger) (*result.Result, error) {
	rows, err := db.Query(ctx, sql, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error running SQL")
	}

	var timing *result.Timing
	ret, err := ParseResult("SQL result", 0, sql, timing, rows, logger)
	if err != nil {
		return nil, errors.Wrap(err, "error constructing result for SQL")
	}

	return ret, nil
}
