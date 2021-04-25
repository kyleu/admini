package database

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/util"

	"github.com/jmoiron/sqlx"
)

// Database access service
type Service struct {
	debug bool
	db    *sqlx.DB
}

// Returns a fresh Service
func NewService(debug bool, db *sqlx.DB) *Service {
	return &Service{debug: debug, db: db}
}

// Begins a transaction, be sure to commit it when you're done
func (s *Service) StartTransaction() (*sqlx.Tx, error) {
	if s.debug {
		util.LogInfo("opening transaction")
	}
	return s.db.Beginx()
}

func errMessage(t string, q string, values []interface{}) string {
	return fmt.Sprintf("error running %v sql [%v] with values [%v]", t, strings.TrimSpace(q), valueStrings(values))
}

func logQuery(msg string, q string, values []interface{}) {
	util.LogInfo(fmt.Sprintf("%v {\n  SQL: %v\n  Values: %v\n}", msg, strings.TrimSpace(q), valueStrings(values)))
}
