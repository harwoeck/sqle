package sqle

import (
	"database/sql"
	"fmt"
)

// Select selects the `query` string from the database. The `args` interface
// slice should contain all primitive value arguments. The `dest` interface
// slice should contain a collection of pointer to primitive types. The results
// of the query will be saved into these pointers.
func (s Sqle) Select(query string, args []interface{}, dest []interface{}) error {
	_, err := s.SelectExists(query, args, dest)
	return err
}

// SelectExists is the same as `Select`, but additionally returns an boolean
// value, whether or not the row existed.
func (s Sqle) SelectExists(query string, args []interface{}, dest []interface{}) (exists bool, err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return false, err
	}

	exists, err = s.SelectExistsTx(tx, query, args, dest)
	if err != nil {
		rlbErr := tx.Rollback()
		if rlbErr == nil {
			return false, err
		}
		return false, fmt.Errorf("sqle.SelectExists: multiple errors occurred: %q followed by tx.Rollback error: %q", err.Error(), rlbErr.Error())
	}

	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return
}

// SelectExistsTx is the same as `SelectExists` but uses the passed transaction
// `tx` to execute the operations.
func (s Sqle) SelectExistsTx(tx *sql.Tx, query string, args []interface{}, dest []interface{}) (exists bool, err error) {
	if len(dest) == 0 {
		return false, fmt.Errorf("sqle.SelectExistsTx: no dest to scan to")
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	if len(args) > 0 {
		err = stmt.QueryRow(args...).Scan(dest...)
	} else if len(args) == 0 {
		err = stmt.QueryRow().Scan(dest...)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// SelectRange selects a range of results from the database, defined by the
// `query` and it's arguments. The `args` interface slice should contain all
// primitive value arguments. The `dest` interface slice should contain a
// collection of pointer to primitive types. The results of the query will be
// saved into these pointers. As soon as one row has been loaded the
// `handleRow` callback will be called. It is the package caller's
// responsibility to copy the values from the `dest` collection into another
// data structure. After returning the `handleRow` function the values of
// `dest` will be overwritten with the next row's values.
func (s Sqle) SelectRange(query string, args []interface{}, dest []interface{}, handleRow func()) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	err = s.SelectRangeTx(tx, query, args, dest, handleRow)
	if err != nil {
		rlbErr := tx.Rollback()
		if rlbErr == nil {
			return err
		}

		return fmt.Errorf("sqle.SelectRange: multiple errors occurred: %q followed by tx.Rollback error: %q", err.Error(), rlbErr.Error())
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// SelectRangeTx is the same as `SelectRange` but uses the passed transaction
// `tx` to execute the operations.
func (s Sqle) SelectRangeTx(tx *sql.Tx, query string, args []interface{}, dest []interface{}, handleRow func()) error {
	if len(dest) == 0 {
		return fmt.Errorf("sqle.SelectRangeTx: no dest to scan to")
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return err
		}
		handleRow()
	}
	if rows.Err() != nil {
		return err
	}

	return nil
}
