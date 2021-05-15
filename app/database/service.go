package database

import (
	"fmt"
	"strings"

	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"
)

// Database access service
type Service struct {
	DatabaseName string
	SchemaName   string
	debug        *zap.SugaredLogger
	db           *sqlx.DB
}

// Returns a fresh Service
func NewService(dbName string, schName string, debug *zap.SugaredLogger, db *sqlx.DB) *Service {
	return &Service{DatabaseName: dbName, SchemaName: schName, debug: debug, db: db}
}

// Begins a transaction, be sure to commit it when you're done
func (s *Service) StartTransaction() (*sqlx.Tx, error) {
	if s.debug != nil {
		s.debug.Info("opening transaction")
	}
	return s.db.Beginx()
}

func errMessage(t string, q string, values []interface{}) string {
	return fmt.Sprintf("error running %v sql [%v] with values [%v]", t, strings.TrimSpace(q), valueStrings(values))
}

func (s *Service) logQuery(msg string, q string, values []interface{}) {
	if s.debug != nil {
		s.debug.Info("%v {\n  SQL: %v\n  Values: %v\n}", msg, strings.TrimSpace(q), valueStrings(values))
	}
}
