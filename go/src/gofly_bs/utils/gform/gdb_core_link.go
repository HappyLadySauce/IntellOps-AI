package gform

import (
	"database/sql"
)

// dbLink is used to implement interface Link for DB.
type dbLink struct {
	*sql.DB         // Underlying DB object.
	isOnMaster bool // isOnMaster marks whether current link is operated on master node.
}

// txLink is used to implement interface Link for TX.
type txLink struct {
	*sql.Tx
}

// IsTransaction returns if current Link is a transaction.
func (l *dbLink) IsTransaction() bool {
	return false
}

// IsOnMaster checks and returns whether current link is operated on master node.
func (l *dbLink) IsOnMaster() bool {
	return l.isOnMaster
}

// IsTransaction returns if current Link is a transaction.
func (l *txLink) IsTransaction() bool {
	return true
}

// IsOnMaster checks and returns whether current link is operated on master node.
// Note that, transaction operation is always operated on master node.
func (l *txLink) IsOnMaster() bool {
	return true
}
