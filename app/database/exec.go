package database

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Runs a SQL insert statement, returning an optional error
func (s *Service) Insert(q string, tx *sqlx.Tx, values ...interface{}) error {
	if s.debug {
		logQuery("inserting row", q, values)
	}
	aff, err := s.execUnknown(q, tx, values...)
	if err != nil {
		return err
	}
	if aff == 0 {
		return fmt.Errorf("no rows affected by insert using sql [%v] and %v values", q, len(values))
	}
	return nil
}

// Runs a SQL update statement, returning the number of affected rows and an optional error
func (s *Service) Update(q string, tx *sqlx.Tx, expected int, values ...interface{}) (int, error) {
	return s.process("updating", "updated", q, tx, expected, values...)
}

// Runs a SQL update statement for a single row, returning an optional error and verifying that a single row was updated
func (s *Service) UpdateOne(q string, tx *sqlx.Tx, values ...interface{}) error {
	_, err := s.Update(q, tx, 1, values...)
	return err
}

// Runs a SQL delete statement, returning the number of affected rows and an optional error
func (s *Service) Delete(q string, tx *sqlx.Tx, expected int, values ...interface{}) (int, error) {
	return s.process("deleting", "deleted", q, tx, expected, values...)
}

// Runs a SQL delete statement for a single row, returning an optional error and verifying that a single row was removed
func (s *Service) DeleteOne(q string, tx *sqlx.Tx, values ...interface{}) error {
	_, err := s.Delete(q, tx, 1, values...)
	if err != nil {
		return fmt.Errorf(errMessage("delete", q, values)+": %w", err)
	}
	return err
}

// Runs an arbitrary SQL statement, returning the number of affected rows and an optional error
func (s *Service) Exec(q string, tx *sqlx.Tx, expected int, values ...interface{}) (int, error) {
	return s.process("executing", "executed", q, tx, expected, values...)
}

func (s *Service) execUnknown(q string, tx *sqlx.Tx, values ...interface{}) (int, error) {
	var err error
	var ret sql.Result
	if tx == nil {
		r, e := s.db.Exec(q, values...)
		ret = r
		err = e
	} else {
		r, e := tx.Exec(q, values...)
		ret = r
		err = e
	}
	if err != nil {
		return 0, fmt.Errorf(errMessage("exec", q, values)+": %w", err)
	}
	aff, _ := ret.RowsAffected()
	// if err != nil {
	// 	return 0, fmt.Errorf("%w", err)
	// }
	return int(aff), nil
}

func (s *Service) process(key string, past string, q string, tx *sqlx.Tx, expected int, values ...interface{}) (int, error) {
	if s.debug {
		logQuery(fmt.Sprintf("%v [%v] rows", key, expected), q, values)
	}

	aff, err := s.execUnknown(q, tx, values...)
	if err != nil {
		return 0, fmt.Errorf(errMessage(past, q, values)+": %w", err)
	}
	if expected > -1 && aff != expected {
		const msg = "expected [%v] %v row(s), but [%v] records affected from sql [%v] with values [%s]"
		return aff, fmt.Errorf(msg, expected, past, aff, q, valueStrings(values))
	}
	return aff, nil
}
