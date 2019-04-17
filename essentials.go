package sqle

import (
	"context"
	"database/sql"
)

// Essentials is a common minimum accros all sqle implementations. It defines
// only functions that can be executed using database/sql.DB from the standard
// library.
type Essentials interface {
	Exec(ctx context.Context, query string, args ...interface{}) error
	ExecTx(ctx context.Context, tx *sql.Tx, query string, args ...interface{}) error
	ExecID(ctx context.Context, query string, args ...interface{}) (lastInsertID int64, err error)
	ExecIDTx(ctx context.Context, tx *sql.Tx, query string, args ...interface{}) (lastInsertID int64, err error)
	ExecAffected(ctx context.Context, query string, args ...interface{}) (rowsAffected int64, err error)
	ExecAffectedTx(ctx context.Context, tx *sql.Tx, query string, args ...interface{}) (rowsAffected int64, err error)
	ExecRes(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	ExecResTx(ctx context.Context, tx *sql.Tx, query string, args ...interface{}) (sql.Result, error)

	Select(ctx context.Context, query string, args []interface{}, dests []interface{}) error
	SelectTx(ctx context.Context, tx *sql.Tx, query string, args []interface{}, dests []interface{}) error
	SelectExists(ctx context.Context, query string, args []interface{}, dests []interface{}) (exists bool, err error)
	SelectExistsTx(ctx context.Context, tx *sql.Tx, query string, args []interface{}, dests []interface{}) (exists bool, err error)
	SelectRange(ctx context.Context, query string, args []interface{}, dests []interface{}, handleRow func()) error
	SelectRangeTx(ctx context.Context, tx *sql.Tx, query string, args []interface{}, dests []interface{}, handleRow func()) error

	UnsafeExecBatch(ctx context.Context, statements []string) error
	UnsafeExecBatchTx(ctx context.Context, tx *sql.Tx, statements []string) error
}
