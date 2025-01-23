package pgsql

import (
	"context"
	"database/sql"

	"gofly/utils/gform"
	"gofly/utils/tools/gcode"
	"gofly/utils/tools/gerror"
)

// DoInsert inserts or updates data forF given table.
func (d *Driver) DoInsert(ctx context.Context, link gform.Link, table string, list gform.List, option gform.DoInsertOption) (result sql.Result, err error) {
	switch option.InsertOption {
	case gform.InsertOptionSave:
		return nil, gerror.NewCode(
			gcode.CodeNotSupported,
			`Save operation is not supported by pgsql driver`,
		)

	case gform.InsertOptionReplace:
		return nil, gerror.NewCode(
			gcode.CodeNotSupported,
			`Replace operation is not supported by pgsql driver`,
		)

	case gform.InsertOptionIgnore:
		return nil, gerror.NewCode(
			gcode.CodeNotSupported,
			`Insert ignore operation is not supported by pgsql driver`,
		)

	case gform.InsertOptionDefault:
		tableFields, err := d.GetCore().GetDB().TableFields(ctx, table)
		if err == nil {
			for _, field := range tableFields {
				if field.Key == "pri" {
					pkField := *field
					ctx = context.WithValue(ctx, internalPrimaryKeyInCtx, pkField)
					break
				}
			}
		}
	}
	return d.Core.DoInsert(ctx, link, table, list, option)
}
