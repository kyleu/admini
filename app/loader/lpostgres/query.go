package lpostgres

import (
	"github.com/kyleu/admini/app/result"
	"github.com/pkg/errors"
)

func (l *Loader) Query(source string, cfg []byte, sql string) (*result.Result, error) {
	db, err := l.openDatabase(source, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database")
	}

	rows, err := db.Query(sql, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error running SQL")
	}

	var timing *result.Timing
	ret, err := ParseResult("SQL result", 0, sql, timing, rows, l.logger)
	if err != nil {
		return nil, errors.Wrap(err, "error constructing result for SQL")
	}

	return ret, nil
}


