package pgsql

import (
	"context"
	"fmt"
	"regexp"

	"gofly/utils/gform"
	"gofly/utils/tools/gregex"
	"gofly/utils/tools/gstr"
	"gofly/utils/tools/gutil"
)

var (
	tablesSqlTmp = `
SELECT
	c.relname
FROM
	pg_class c
INNER JOIN pg_namespace n ON
	c.relnamespace = n.oid
WHERE
	n.nspname = '%s'
	AND c.relkind IN ('r', 'p')
	%s
ORDER BY
	c.relname
`
)

func init() {
	tablesSqlTmp = formatSqlTmp(tablesSqlTmp)
}

// Tables retrieves and returns the tables of current schema.
// It's mainly used in cli tool chain for automatically generating the models.
func (d *Driver) Tables(ctx context.Context, schema ...string) (tables []string, err error) {
	var (
		result     gform.Result
		usedSchema = gutil.GetOrDefaultStr(d.GetConfig().Namespace, schema...)
	)
	if usedSchema == "" {
		usedSchema = defaultSchema
	}
	// DO NOT use `usedSchema` as parameter for function `SlaveLink`.
	link, err := d.SlaveLink(schema...)
	if err != nil {
		return nil, err
	}

	useRelpartbound := ""
	if gstr.CompareVersion(d.version(ctx, link), "10") >= 0 {
		useRelpartbound = "AND c.relpartbound IS NULL"
	}

	var query = fmt.Sprintf(
		tablesSqlTmp,
		usedSchema,
		useRelpartbound,
	)

	query, _ = gregex.ReplaceString(`[\n\r\s]+`, " ", gstr.Trim(query))
	result, err = d.DoSelect(ctx, link, query)
	if err != nil {
		return
	}
	for _, m := range result {
		for _, v := range m {
			tables = append(tables, v.String())
		}
	}
	return
}

// version checks and returns the database version.
func (d *Driver) version(ctx context.Context, link gform.Link) string {
	result, err := d.DoSelect(ctx, link, "SELECT version();")
	if err != nil {
		return ""
	}
	if len(result) > 0 {
		if v, ok := result[0]["version"]; ok {
			matches := regexp.MustCompile(`PostgreSQL (\d+\.\d+)`).FindStringSubmatch(v.String())
			if len(matches) >= 2 {
				return matches[1]
			}
		}
	}
	return ""
}
