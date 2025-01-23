package gform

import (
	"context"
)

// Transaction wraps the transaction logic using function `f`.
// It rollbacks the transaction and returns the error from function `f` if
// it returns non-nil error. It commits the transaction and returns nil if
// function `f` returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function `f`
// as it is automatically handled by this function.
func (m *Model) Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) (err error) {
	if ctx == nil {
		ctx = m.GetCtx()
	}
	if m.tx != nil {
		return m.tx.Transaction(ctx, f)
	}
	return m.db.Transaction(ctx, f)
}
