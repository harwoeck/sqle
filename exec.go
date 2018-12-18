package sqle

import (
	"database/sql"
	"fmt"
)

// Exec will execute the query with it's argument against the database and
// returns any occurring errors.
func (s Sqle) Exec(query string, args ...interface{}) error {
	_, err := s.ExecID(query, args...)
	return err
}

// ExecID executes the query with it's arguments against the database and
// returns any occurring errors.
//
// If the executed query is an `INSERT` operation with an auto-generated ID
// (for example chosen by the database because of a `AUTO-INCREMENT` flag),
// then `lastInsertID` will contain the generated ID value.
func (s Sqle) ExecID(query string, args ...interface{}) (lastInsertID int64, err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}

	lastInsertID, err = s.ExecTxID(tx, query, args...)
	if err != nil {
		rlbErr := tx.Rollback()
		if rlbErr == nil {
			return 0, err
		}
		return 0, fmt.Errorf("sqle.ExecID: multiple errors occurred: %q followed by tx.Rollback error: %q", err.Error(), rlbErr.Error())
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return
}

// ExecTx is the same as `Exec` but uses the passed transaction `tx` to execute
// the operations.
func (s Sqle) ExecTx(tx *sql.Tx, query string, args ...interface{}) error {
	_, err := s.ExecTxID(tx, query, args...)
	return err
}

// ExecTxID is the same as `ExecID` but uses the passed transaction `tx` to
// execute the operations.
func (s Sqle) ExecTxID(tx *sql.Tx, query string, args ...interface{}) (lastInsertID int64, err error) {
	stmt, err := tx.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}

	lastInsertID, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

// ExecBatch runs 'Exec` without any arguments for every entry in the slice
// parameter.
//
// This method IS NOT SAFE AGAINST SQL-INJECTION. Use it only with trusted
// input!
func (s Sqle) ExecBatch(queries []string) error {
	for _, q := range queries {
		err := s.Exec(q)
		if err != nil {
			return err
		}
	}

	return nil
}
