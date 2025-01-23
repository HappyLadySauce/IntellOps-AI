package pgsql

import (
	_ "github.com/lib/pq"

	"gofly/utils/gform"
	"gofly/utils/tools/gctx"
	"gofly/utils/tools/gregex"
	"gofly/utils/tools/gstr"
)

// Driver is the driver for postgresql database.
type Driver struct {
	*gform.Core
}

const (
	internalPrimaryKeyInCtx gctx.StrKey = "primary_key"
	defaultSchema           string      = "public"
	quoteChar               string      = `"`
)

func init() {
	if err := gform.Register(`pgsql`, New()); err != nil {
		panic(err)
	}
}

// formatSqlTmp formats sql template string into one line.
func formatSqlTmp(sqlTmp string) string {
	var err error
	// format sql template string.
	sqlTmp, err = gregex.ReplaceString(`[\n\r\s]+`, " ", gstr.Trim(sqlTmp))
	if err != nil {
		panic(err)
	}
	sqlTmp, err = gregex.ReplaceString(`\s{2,}`, " ", gstr.Trim(sqlTmp))
	if err != nil {
		panic(err)
	}
	return sqlTmp
}

// New create and returns a driver that implements gform.Driver, which supports operations for PostgreSql.
func New() gform.Driver {
	return &Driver{}
}

// New creates and returns a database object for postgresql.
// It implements the interface of gform.Driver for extra database driver installation.
func (d *Driver) New(core *gform.Core, node *gform.ConfigNode) (gform.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

// GetChars returns the security char for this type of database.
func (d *Driver) GetChars() (charLeft string, charRight string) {
	return quoteChar, quoteChar
}
