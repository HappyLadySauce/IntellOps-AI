package pgsql

import (
	"context"
	"fmt"

	"gofly/utils/gform"
	"gofly/utils/tools/gregex"
)

// DoFilter deals with the sql string before commits it to underlying sql driver.
func (d *Driver) DoFilter(
	ctx context.Context, link gform.Link, sql string, args []interface{},
) (newSql string, newArgs []interface{}, err error) {
	var index int
	// Convert placeholder char '?' to string "$x".
	newSql, err = gregex.ReplaceStringFunc(`\?`, sql, func(s string) string {
		index++
		return fmt.Sprintf(`$%d`, index)
	})
	if err != nil {
		return "", nil, err
	}
	// Handle pgsql jsonb feature support, which contains place-holder char '?'.
	// Refer:
	// https://github.com/gogf/gf/issues/1537
	// https://www.postgresql.org/docs/12/functions-json.html
	newSql, err = gregex.ReplaceStringFuncMatch(`(::jsonb([^\w\d]*)\$\d)`, newSql, func(match []string) string {
		return fmt.Sprintf(`::jsonb%s?`, match[2])
	})
	if err != nil {
		return "", nil, err
	}
	newSql, err = gregex.ReplaceString(` LIMIT (\d+),\s*(\d+)`, ` LIMIT $2 OFFSET $1`, newSql)
	if err != nil {
		return "", nil, err
	}
	newArgs = args
	return d.Core.DoFilter(ctx, link, newSql, newArgs)
}
