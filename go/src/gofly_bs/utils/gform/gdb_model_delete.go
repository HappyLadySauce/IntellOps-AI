package gform

import (
	"database/sql"

	"gofly/utils/tools/gcode"
	"gofly/utils/tools/gerror"
	"gofly/utils/tools/gstr"
	"gofly/utils/tools/intlog"
)

// Delete does "DELETE FROM ... " statement for the model.
// The optional parameter `where` is the same as the parameter of Model.Where function,
// see Model.Where.
func (m *Model) Delete(where ...interface{}) (result sql.Result, err error) {
	var ctx = m.GetCtx()
	if len(where) > 0 {
		return m.Where(where[0], where[1:]...).Delete()
	}
	defer func() {
		if err == nil {
			m.checkAndRemoveSelectCache(ctx)
		}
	}()
	var (
		conditionWhere, conditionExtra, conditionArgs = m.formatCondition(ctx, false, false)
		conditionStr                                  = conditionWhere + conditionExtra
		fieldNameDelete, fieldTypeDelete              = m.softTimeMaintainer().GetFieldNameAndTypeForDelete(
			ctx, "", m.tablesInit,
		)
	)
	if m.unscoped {
		fieldNameDelete = ""
	}
	if !gstr.ContainsI(conditionStr, " WHERE ") || (fieldNameDelete != "" && !gstr.ContainsI(conditionStr, " AND ")) {
		intlog.Printf(
			ctx,
			`sql condition string "%s" has no WHERE for DELETE operation, fieldNameDelete: %s`,
			conditionStr, fieldNameDelete,
		)
		return nil, gerror.NewCode(
			gcode.CodeMissingParameter,
			"there should be WHERE condition statement for DELETE operation",
		)
	}

	// Soft deleting.
	if fieldNameDelete != "" {
		dataHolder, dataValue := m.softTimeMaintainer().GetDataByFieldNameAndTypeForDelete(
			ctx, "", fieldNameDelete, fieldTypeDelete,
		)
		in := &HookUpdateInput{
			internalParamHookUpdate: internalParamHookUpdate{
				internalParamHook: internalParamHook{
					link: m.getLink(true),
				},
				handler: m.hookHandler.Update,
			},
			Model:     m,
			Table:     m.tables,
			Data:      dataHolder,
			Condition: conditionStr,
			Args:      append([]interface{}{dataValue}, conditionArgs...),
		}
		return in.Next(ctx)
	}

	in := &HookDeleteInput{
		internalParamHookDelete: internalParamHookDelete{
			internalParamHook: internalParamHook{
				link: m.getLink(true),
			},
			handler: m.hookHandler.Delete,
		},
		Model:     m,
		Table:     m.tables,
		Condition: conditionStr,
		Args:      conditionArgs,
	}
	return in.Next(ctx)
}
