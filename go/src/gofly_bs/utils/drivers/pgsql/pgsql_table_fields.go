package pgsql

import (
	"context"
	"fmt"

	"gofly/utils/gform"
	"gofly/utils/tools/gutil"
)

var (
	tableFieldsSqlTmp = `
SELECT a.attname AS field, t.typname AS type,a.attnotnull as null,
    (case when d.contype is not null then 'pri' else '' end)  as key
      ,ic.column_default as default_value,b.description as comment
      ,coalesce(character_maximum_length, numeric_precision, -1) as length
      ,numeric_scale as scale
FROM pg_attribute a
         left join pg_class c on a.attrelid = c.oid
         left join pg_constraint d on d.conrelid = c.oid and a.attnum = d.conkey[1]
         left join pg_description b ON a.attrelid=b.objoid AND a.attnum = b.objsubid
         left join pg_type t ON a.atttypid = t.oid
         left join information_schema.columns ic on ic.column_name = a.attname and ic.table_name = c.relname
WHERE c.relname = '%s' and a.attisdropped is false and a.attnum > 0
ORDER BY a.attnum`
)

func init() {
	tableFieldsSqlTmp = formatSqlTmp(tableFieldsSqlTmp)
}

// TableFields retrieves and returns the fields' information of specified table of current schema.
func (d *Driver) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*gform.TableField, err error) {
	var (
		result     gform.Result
		link       gform.Link
		usedSchema = gutil.GetOrDefaultStr(d.GetSchema(), schema...)
		// TODO duplicated `id` result?
		structureSql = fmt.Sprintf(tableFieldsSqlTmp, table)
	)
	if link, err = d.SlaveLink(usedSchema); err != nil {
		return nil, err
	}
	result, err = d.DoSelect(ctx, link, structureSql)
	if err != nil {
		return nil, err
	}
	fields = make(map[string]*gform.TableField)
	var (
		index = 0
		name  string
		ok    bool
	)
	for _, m := range result {
		name = m["field"].String()
		// Filter duplicated fields.
		if _, ok = fields[name]; ok {
			continue
		}
		fields[name] = &gform.TableField{
			Index:   index,
			Name:    name,
			Type:    m["type"].String(),
			Null:    !m["null"].Bool(),
			Key:     m["key"].String(),
			Default: m["default_value"].Val(),
			Comment: m["comment"].String(),
		}
		index++
	}
	return fields, nil
}
