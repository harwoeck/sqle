package sqle

import (
	"context"
	"fmt"
)

// MySQL extends sqle.Std with MySQL specific functions
type MySQL struct {
	Std *Std
}

// UnsafeExists checks whether the statement defined by the `query` and `args`
// would return a result.
//
// This method IS NOT SAFE AGAINST SQL-INJECTION. Use it only with trusted
// input!
func (s *MySQL) UnsafeExists(ctx context.Context, query string, args ...interface{}) (exists bool, err error) {
	query = fmt.Sprintf("SELECT EXISTS (%s)", query)

	err = s.Std.Select(ctx, query, args, []interface{}{&exists})
	return
}

// UnsafeCount counts the rows for a single column in a specified table.
//
// This method IS NOT SAFE AGAINST SQL-INJECTION. Use it only with trusted
// input!
func (s *MySQL) UnsafeCount(ctx context.Context, table, column string) (count int64, err error) {
	query := fmt.Sprintf("SELECT COUNT(%s) FROM %s", column, table)

	err = s.Std.Select(ctx, query, []interface{}{}, []interface{}{&count})
	return
}
