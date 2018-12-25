package sqle

import (
	"context"
	"database/sql"
)

// StdNoCtx wraps `Std` and uses `context.Background()` for all calls to it.
type StdNoCtx struct {
	std *Std
}

// WithContext returns the internal `Std` instance used.
func (no *StdNoCtx) WithContext() Essentials {
	return no.std
}

// Exec calls `Std.Exec` with `context.Background()`
func (no *StdNoCtx) Exec(query string, args ...interface{}) error {
	return no.std.Exec(context.Background(), query, args...)
}

// ExecTx calls `Std.ExecTx` with `context.Background()`
func (no *StdNoCtx) ExecTx(tx *sql.Tx, query string, args ...interface{}) error {
	return no.std.ExecTx(context.Background(), tx, query, args...)
}

// ExecID calls `Std.ExecID` with `context.Background()`
func (no *StdNoCtx) ExecID(query string, args ...interface{}) (lastInsertID int64, err error) {
	return no.std.ExecID(context.Background(), query, args...)
}

// ExecIDTx calls `Std.ExecIDTx` with `context.Background()`
func (no *StdNoCtx) ExecIDTx(tx *sql.Tx, query string, args ...interface{}) (lastInsertID int64, err error) {
	return no.std.ExecIDTx(context.Background(), tx, query, args...)
}

// ExecAffected calls `Std.ExecAffected` with `context.Background()`
func (no *StdNoCtx) ExecAffected(query string, args ...interface{}) (rowsAffected int64, err error) {
	return no.std.ExecAffected(context.Background(), query, args...)
}

// ExecAffectedTx calls `Std.ExecAffectedTx` with `context.Background()`
func (no *StdNoCtx) ExecAffectedTx(tx *sql.Tx, query string, args ...interface{}) (rowsAffected int64, err error) {
	return no.std.ExecAffectedTx(context.Background(), tx, query, args...)
}

// ExecRes calls `Std.ExecRes` with `context.Background()`
func (no *StdNoCtx) ExecRes(query string, args ...interface{}) (sql.Result, error) {
	return no.std.ExecRes(context.Background(), query, args...)
}

// ExecResTx calls `Std.ExecResTx` with `context.Background()`
func (no *StdNoCtx) ExecResTx(tx *sql.Tx, query string, args ...interface{}) (sql.Result, error) {
	return no.std.ExecResTx(context.Background(), tx, query, args...)
}

// Select calls `Std.Select` with `context.Background()`
func (no *StdNoCtx) Select(query string, args []interface{}, dests []interface{}) error {
	return no.std.Select(context.Background(), query, args, dests)
}

// SelectTx calls `Std.SelectTx` with `context.Background()`
func (no *StdNoCtx) SelectTx(tx *sql.Tx, query string, args []interface{}, dests []interface{}) error {
	return no.std.SelectTx(context.Background(), tx, query, args, dests)
}

// SelectExists calls `Std.SelectExists` with `context.Background()`
func (no *StdNoCtx) SelectExists(query string, args []interface{}, dests []interface{}) (exists bool, err error) {
	return no.std.SelectExists(context.Background(), query, args, dests)
}

// SelectExistsTx calls `Std.SelectExistsTx` with `context.Background()`
func (no *StdNoCtx) SelectExistsTx(tx *sql.Tx, query string, args []interface{}, dests []interface{}) (exists bool, err error) {
	return no.std.SelectExistsTx(context.Background(), tx, query, args, dests)
}

// SelectRange calls `Std.SelectRange` with `context.Background()`
func (no *StdNoCtx) SelectRange(query string, args []interface{}, dests []interface{}, handleRow func()) error {
	return no.std.SelectRange(context.Background(), query, args, dests, handleRow)
}

// SelectRangeTx calls `Std.SelectRangeTx` with `context.Background()`
func (no *StdNoCtx) SelectRangeTx(tx *sql.Tx, query string, args []interface{}, dests []interface{}, handleRow func()) error {
	return no.std.SelectRangeTx(context.Background(), tx, query, args, dests, handleRow)
}

// UnsafeExecBatch calls `Std.UnsafeExecBatch` with `context.Background()`
func (no *StdNoCtx) UnsafeExecBatch(statements []string) error {
	return no.std.UnsafeExecBatch(context.Background(), statements)

}

// UnsafeExecBatchTx calls `Std.UnsafeExecBatchTx` with `context.Background()`
func (no *StdNoCtx) UnsafeExecBatchTx(tx *sql.Tx, statements []string) error {
	return no.std.UnsafeExecBatchTx(context.Background(), tx, statements)
}
