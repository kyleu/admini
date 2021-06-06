package database

import (
	"go.uber.org/zap"

	"github.com/pkg/errors"

	// load sqlite driver.
	_ "modernc.org/sqlite"

	"github.com/jmoiron/sqlx"
)

type SQLiteParams struct {
	File   string `json:"file"`
	Schema string `json:"schema,omitempty"`
	Debug  bool   `json:"debug,omitempty"`
}

func OpenSQLiteDatabase(params *SQLiteParams, logger *zap.SugaredLogger) (*Service, error) {
	if params.File == "" {
		return nil, errors.New("need filename for SQLite database")
	}
	db, err := sqlx.Open("sqlite", params.File)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database")
	}

	var log *zap.SugaredLogger
	if params.Debug {
		log = logger
	}

	svc := NewService(params.File, params.Schema, log, db)

	return svc, nil
}
