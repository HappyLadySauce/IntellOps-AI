package gform

import (
	"context"
	"fmt"

	"gofly/utils/tools/garray"
	"gofly/utils/tools/gcache"
	"gofly/utils/tools/gcode"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gerror"
	"gofly/utils/tools/gregex"
	"gofly/utils/tools/gstr"
	"gofly/utils/tools/gtime"
	"gofly/utils/tools/gutil"
	"gofly/utils/tools/intlog"
)

// SoftTimeType custom defines the soft time field type.
type SoftTimeType int

const (
	SoftTimeTypeAuto           SoftTimeType = 0 // (Default)Auto detect the field type by table field type.
	SoftTimeTypeTime           SoftTimeType = 1 // Using datetime as the field value.
	SoftTimeTypeTimestamp      SoftTimeType = 2 // In unix seconds.
	SoftTimeTypeTimestampMilli SoftTimeType = 3 // In unix milliseconds.
	SoftTimeTypeTimestampMicro SoftTimeType = 4 // In unix microseconds.
	SoftTimeTypeTimestampNano  SoftTimeType = 5 // In unix nanoseconds.
)

// SoftTimeOption is the option to customize soft time feature for Model.
type SoftTimeOption struct {
	SoftTimeType SoftTimeType // The value type for soft time field.
}

type softTimeMaintainer struct {
	*Model
}

type iSoftTimeMaintainer interface {
	GetFieldNameAndTypeForCreate(
		ctx context.Context, schema string, table string,
	) (fieldName string, fieldType LocalType)

	GetFieldNameAndTypeForUpdate(
		ctx context.Context, schema string, table string,
	) (fieldName string, fieldType LocalType)

	GetFieldNameAndTypeForDelete(
		ctx context.Context, schema string, table string,
	) (fieldName string, fieldType LocalType)

	GetValueByFieldTypeForCreateOrUpdate(
		ctx context.Context, fieldType LocalType, isDeletedField bool,
	) (dataValue any)

	GetDataByFieldNameAndTypeForDelete(
		ctx context.Context, fieldPrefix, fieldName string, fieldType LocalType,
	) (dataHolder string, dataValue any)

	GetWhereConditionForDelete(ctx context.Context) string
}

// getSoftFieldNameAndTypeCacheItem is the internal struct for storing create/update/delete fields.
type getSoftFieldNameAndTypeCacheItem struct {
	FieldName string
	FieldType LocalType
}

var (
	// Default field names of table for automatic-filled for record creating.
	createdFieldNames = "created_at"
	// Default field names of table for automatic-filled for record updating.
	updatedFieldNames = "updated_at"
	// Default field names of table for automatic-filled for record deleting.
	deletedFieldNames = "deleted_at"
)

// SoftTime sets the SoftTimeOption to customize soft time feature for Model.
func (m *Model) SoftTime(option SoftTimeOption) *Model {
	model := m.getModel()
	model.softTimeOption = option
	return model
}

// Unscoped disables the soft time feature for insert, update and delete operations.
func (m *Model) Unscoped() *Model {
	model := m.getModel()
	model.unscoped = true
	return model
}

func (m *Model) softTimeMaintainer() iSoftTimeMaintainer {
	return &softTimeMaintainer{
		m,
	}
}

// 创建时间
// GetFieldNameAndTypeForCreate checks and returns the field name for record creating time.
// If there's no field name for storing creating time, it returns an empty string.
// It checks the key with or without cases or chars '-'/'_'/'.'/' '.
func (m *softTimeMaintainer) GetFieldNameAndTypeForCreate(
	ctx context.Context, schema string, table string,
) (fieldName string, fieldType LocalType) {
	// It checks whether this feature disabled.
	if m.db.GetConfig().TimeMaintainDisabled {
		return "", LocalTypeUndefined
	}
	tableName := ""
	if table != "" {
		tableName = table
	} else {
		tableName = m.tablesInit
	}
	config := m.db.GetConfig()
	if config.CreatedAt != "" {
		return m.getSoftFieldNameAndType(
			ctx, schema, tableName, config.CreatedAt,
		)
	}
	return m.getSoftFieldNameAndType(
		ctx, schema, tableName, createdFieldNames,
	)
}

// 更新时间
// GetFieldNameAndTypeForUpdate checks and returns the field name for record updating time.
// If there's no field name for storing updating time, it returns an empty string.
// It checks the key with or without cases or chars '-'/'_'/'.'/' '.
func (m *softTimeMaintainer) GetFieldNameAndTypeForUpdate(
	ctx context.Context, schema string, table string,
) (fieldName string, fieldType LocalType) {
	// It checks whether this feature disabled.
	if m.db.GetConfig().TimeMaintainDisabled {
		return "", LocalTypeUndefined
	}
	tableName := ""
	if table != "" {
		tableName = table
	} else {
		tableName = m.tablesInit
	}
	config := m.db.GetConfig()
	if config.UpdatedAt != "" {
		return m.getSoftFieldNameAndType(
			ctx, schema, tableName, config.UpdatedAt,
		)
	}
	return m.getSoftFieldNameAndType(
		ctx, schema, tableName, updatedFieldNames,
	)
}

// 删除时间
// GetFieldNameAndTypeForDelete checks and returns the field name for record deleting time.
// If there's no field name for storing deleting time, it returns an empty string.
// It checks the key with or without cases or chars '-'/'_'/'.'/' '.
func (m *softTimeMaintainer) GetFieldNameAndTypeForDelete(
	ctx context.Context, schema string, table string,
) (fieldName string, fieldType LocalType) {
	// It checks whether this feature disabled.
	if m.db.GetConfig().TimeMaintainDisabled {
		return "", LocalTypeUndefined
	}
	tableName := ""
	if table != "" {
		tableName = table
	} else {
		tableName = m.tablesInit
	}
	config := m.db.GetConfig()
	if config.DeletedAt != "" {
		return m.getSoftFieldNameAndType(
			ctx, schema, tableName, config.DeletedAt,
		)
	}
	return m.getSoftFieldNameAndType(
		ctx, schema, tableName, deletedFieldNames,
	)
}

// getSoftFieldName retrieves and returns the field name of the table for possible key.
func (m *softTimeMaintainer) getSoftFieldNameAndType(
	ctx context.Context,
	schema string, table string, checkFiledNames string,
) (fieldName string, fieldType LocalType) {
	var (
		cacheKey      = fmt.Sprintf(`getSoftFieldNameAndType:%s#%s#%v`, schema, table, checkFiledNames)
		cacheDuration = gcache.DurationNoExpire //缓存时间为永久
		cacheFunc     = func(ctx context.Context) (value interface{}, err error) {
			fieldsMap, _ := m.TableFields(table, schema)
			// m.db.GetCore().logger.Debug(ctx, "执行函数了88", table, schema, fieldsMap)
			if len(fieldsMap) > 0 {
				fieldName, _ = gutil.MapPossibleItemByKey(
					gconv.Map(fieldsMap), checkFiledNames,
				)
				if fieldName != "" {
					fieldType, _ = m.db.CheckLocalTypeForField(
						ctx, fieldsMap[fieldName].Type, nil,
					)
					var cacheItem = getSoftFieldNameAndTypeCacheItem{
						FieldName: fieldName,
						FieldType: fieldType,
					}
					// m.db.GetCore().logger.Debug(ctx, "执行函数了88-获取数据", cacheItem)
					return cacheItem, nil
				}
			}
			return nil, nil
		}
	)
	// m.db.GetCore().logger.Debug(ctx, "缓存key：", cacheKey)
	result, err := m.db.GetCache().GetOrSetFunc(ctx, cacheKey, cacheFunc, cacheDuration)
	// m.db.GetCore().logger.Debug(ctx, "从缓存中获取：", result, err)
	if err != nil {
		intlog.Error(ctx, err)
	}
	if result != nil {
		var cacheItem = result.Val().(getSoftFieldNameAndTypeCacheItem)
		fieldName = cacheItem.FieldName
		fieldType = cacheItem.FieldType
	} else {
		fieldName = ""
		fieldType = ""
	}
	return
}

// GetWhereConditionForDelete retrieves and returns the condition string for soft deleting.
// It supports multiple tables string like:
// "user u, user_detail ud"
// "user u LEFT JOIN user_detail ud ON(ud.uid=u.uid)"
// "user LEFT JOIN user_detail ON(user_detail.uid=user.uid)"
// "user u LEFT JOIN user_detail ud ON(ud.uid=u.uid) LEFT JOIN user_stats us ON(us.uid=u.uid)".
func (m *softTimeMaintainer) GetWhereConditionForDelete(ctx context.Context) string {
	if m.unscoped {
		return ""
	}
	conditionArray := garray.NewStrArray()
	if gstr.Contains(m.tables, " JOIN ") {
		// Base table.
		tableMatch, _ := gregex.MatchString(`(.+?) [A-Z]+ JOIN`, m.tables)
		conditionArray.Append(m.getConditionOfTableStringForSoftDeleting(ctx, tableMatch[1]))
		// Multiple joined tables, exclude the sub query sql which contains char '(' and ')'.
		tableMatches, _ := gregex.MatchAllString(`JOIN ([^()]+?) ON`, m.tables)
		for _, match := range tableMatches {
			conditionArray.Append(m.getConditionOfTableStringForSoftDeleting(ctx, match[1]))
		}
	}
	if conditionArray.Len() == 0 && gstr.Contains(m.tables, ",") {
		// Multiple base tables.
		for _, s := range gstr.SplitAndTrim(m.tables, ",") {
			conditionArray.Append(m.getConditionOfTableStringForSoftDeleting(ctx, s))
		}
	}
	conditionArray.FilterEmpty()
	if conditionArray.Len() > 0 {
		return conditionArray.Join(" AND ")
	}
	// Only one table.
	fieldName, fieldType := m.GetFieldNameAndTypeForDelete(ctx, "", m.tablesInit)
	if fieldName != "" {
		return m.getConditionByFieldNameAndTypeForSoftDeleting(ctx, "", fieldName, fieldType)
	}
	return ""
}

// getConditionOfTableStringForSoftDeleting does something as its name describes.
// Examples for `s`:
// - `test`.`demo` as b
// - `test`.`demo` b
// - `demo`
// - demo
func (m *softTimeMaintainer) getConditionOfTableStringForSoftDeleting(ctx context.Context, s string) string {
	var (
		table  string
		schema string
		array1 = gstr.SplitAndTrim(s, " ")
		array2 = gstr.SplitAndTrim(array1[0], ".")
	)
	if len(array2) >= 2 {
		table = array2[1]
		schema = array2[0]
	} else {
		table = array2[0]
	}
	fieldName, fieldType := m.GetFieldNameAndTypeForDelete(ctx, schema, table)
	if fieldName == "" {
		return ""
	}
	if len(array1) >= 3 {
		return m.getConditionByFieldNameAndTypeForSoftDeleting(ctx, array1[2], fieldName, fieldType)
	}
	if len(array1) >= 2 {
		return m.getConditionByFieldNameAndTypeForSoftDeleting(ctx, array1[1], fieldName, fieldType)
	}
	return m.getConditionByFieldNameAndTypeForSoftDeleting(ctx, table, fieldName, fieldType)
}

// GetDataByFieldNameAndTypeForDelete creates and returns the placeholder and value for
// specified field name and type in soft-deleting scenario.
func (m *softTimeMaintainer) GetDataByFieldNameAndTypeForDelete(
	ctx context.Context, fieldPrefix, fieldName string, fieldType LocalType,
) (dataHolder string, dataValue any) {
	var (
		quotedFieldPrefix = m.db.GetCore().QuoteWord(fieldPrefix)
		quotedFieldName   = m.db.GetCore().QuoteWord(fieldName)
	)
	if quotedFieldPrefix != "" {
		quotedFieldName = fmt.Sprintf(`%s.%s`, quotedFieldPrefix, quotedFieldName)
	}
	dataHolder = fmt.Sprintf(`%s=?`, quotedFieldName)
	dataValue = m.GetValueByFieldTypeForCreateOrUpdate(ctx, fieldType, false)
	return
}

func (m *softTimeMaintainer) getConditionByFieldNameAndTypeForSoftDeleting(
	ctx context.Context, fieldPrefix, fieldName string, fieldType LocalType,
) string {
	var (
		quotedFieldPrefix = m.db.GetCore().QuoteWord(fieldPrefix)
		quotedFieldName   = m.db.GetCore().QuoteWord(fieldName)
	)
	if quotedFieldPrefix != "" {
		quotedFieldName = fmt.Sprintf(`%s.%s`, quotedFieldPrefix, quotedFieldName)
	}
	switch m.softTimeOption.SoftTimeType {
	case SoftTimeTypeAuto:
		switch fieldType {
		case LocalTypeDate, LocalTypeDatetime:
			return fmt.Sprintf(`%s IS NULL`, quotedFieldName)
		case LocalTypeInt, LocalTypeUint, LocalTypeInt64, LocalTypeBool:
			return fmt.Sprintf(`%s=0`, quotedFieldName)
		default:
			intlog.Errorf(
				ctx,
				`invalid field type "%s" of field name "%s" with prefix "%s" for soft deleting condition`,
				fieldType, fieldName, fieldPrefix,
			)
		}

	case SoftTimeTypeTime:
		return fmt.Sprintf(`%s IS NULL`, quotedFieldName)

	default:
		return fmt.Sprintf(`%s=0`, quotedFieldName)
	}
	return ""
}

// GetValueByFieldTypeForCreateOrUpdate creates and returns the value for specified field type,
// usually for creating or updating operations.
func (m *softTimeMaintainer) GetValueByFieldTypeForCreateOrUpdate(
	ctx context.Context, fieldType LocalType, isDeletedField bool,
) any {
	var value any
	if isDeletedField {
		switch fieldType {
		case LocalTypeDate, LocalTypeDatetime:
			value = nil
		default:
			value = 0
		}
		return value
	}
	switch m.softTimeOption.SoftTimeType {
	case SoftTimeTypeAuto:
		switch fieldType {
		case LocalTypeDate, LocalTypeDatetime:
			value = gtime.Now()
		case LocalTypeInt, LocalTypeUint, LocalTypeInt64:
			value = gtime.Timestamp()
		case LocalTypeBool:
			value = 1
		default:
			intlog.Errorf(
				ctx,
				`invalid field type "%s" for soft deleting data`,
				fieldType,
			)
		}

	default:
		switch fieldType {
		case LocalTypeBool:
			value = 1
		default:
			value = m.createValueBySoftTimeOption(isDeletedField)
		}
	}
	return value
}

func (m *softTimeMaintainer) createValueBySoftTimeOption(isDeletedField bool) any {
	var value any
	if isDeletedField {
		switch m.softTimeOption.SoftTimeType {
		case SoftTimeTypeTime:
			value = nil
		default:
			value = 0
		}
		return value
	}
	switch m.softTimeOption.SoftTimeType {
	case SoftTimeTypeTime:
		value = gtime.Now()
	case SoftTimeTypeTimestamp:
		value = gtime.Timestamp()
	case SoftTimeTypeTimestampMilli:
		value = gtime.TimestampMilli()
	case SoftTimeTypeTimestampMicro:
		value = gtime.TimestampMicro()
	case SoftTimeTypeTimestampNano:
		value = gtime.TimestampNano()
	default:
		panic(gerror.NewCodef(
			gcode.CodeInternalPanic,
			`unrecognized SoftTimeType "%d"`, m.softTimeOption.SoftTimeType,
		))
	}
	return value
}
