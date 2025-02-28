package pgsql

import (
	"context"
	"strings"

	"gofly/utils/gform"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gregex"
	"gofly/utils/tools/gstr"
)

// CheckLocalTypeForField checks and returns corresponding local golang type for given db type.
func (d *Driver) CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{}) (gform.LocalType, error) {
	var typeName string
	match, _ := gregex.MatchString(`(.+?)\((.+)\)`, fieldType)
	if len(match) == 3 {
		typeName = gstr.Trim(match[1])
	} else {
		typeName = fieldType
	}
	typeName = strings.ToLower(typeName)
	switch typeName {
	case
		// For pgsql, int2 = smallint.
		"int2",
		// For pgsql, int4 = integer
		"int4":
		return gform.LocalTypeInt, nil

	case
		// For pgsql, int8 = bigint
		"int8":
		return gform.LocalTypeInt64, nil

	case
		"_int2",
		"_int4":
		return gform.LocalTypeIntSlice, nil

	case
		"_int8":
		return gform.LocalTypeInt64Slice, nil

	default:
		return d.Core.CheckLocalTypeForField(ctx, fieldType, fieldValue)
	}
}

// ConvertValueForLocal converts value to local Golang type of value according field type name from database.
// The parameter `fieldType` is in lower case, like:
// `float(5,2)`, `unsigned double(5,2)`, `decimal(10,2)`, `char(45)`, `varchar(100)`, etc.
func (d *Driver) ConvertValueForLocal(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) {
	typeName, _ := gregex.ReplaceString(`\(.+\)`, "", fieldType)
	typeName = strings.ToLower(typeName)
	switch typeName {
	// For pgsql, int2 = smallint and int4 = integer.
	case "int2", "int4":
		return gconv.Int(gconv.String(fieldValue)), nil

	// For pgsql, int8 = bigint.
	case "int8":
		return gconv.Int64(gconv.String(fieldValue)), nil

	// Int32 slice.
	case
		"_int2", "_int4":
		return gconv.Ints(
			gstr.ReplaceByMap(gconv.String(fieldValue),
				map[string]string{
					"{": "[",
					"}": "]",
				},
			),
		), nil

	// Int64 slice.
	case
		"_int8":
		return gconv.Int64s(
			gstr.ReplaceByMap(gconv.String(fieldValue),
				map[string]string{
					"{": "[",
					"}": "]",
				},
			),
		), nil

	default:
		return d.Core.ConvertValueForLocal(ctx, fieldType, fieldValue)
	}
}
