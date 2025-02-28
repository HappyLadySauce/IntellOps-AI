package pgsql

import (
	"context"
	"database/sql"
	"strings"

	"gofly/utils/gform"
	"gofly/utils/tools/gcode"
	"gofly/utils/tools/gerror"
)

// DoExec commits the sql string and its arguments to underlying driver
// through given link object and returns the execution result.
func (d *Driver) DoExec(ctx context.Context, link gform.Link, sql string, args ...interface{}) (result sql.Result, err error) {
	var (
		isUseCoreDoExec bool   = false // Check whether the default method needs to be used
		primaryKey      string = ""
		pkField         gform.TableField
	)

	// Transaction checks.
	if link == nil {
		if tx := gform.TXFromCtx(ctx, d.GetGroup()); tx != nil {
			// Firstly, check and retrieve transaction link from context.
			link = tx
		} else if link, err = d.MasterLink(); err != nil {
			// Or else it creates one from master node.
			return nil, err
		}
	} else if !link.IsTransaction() {
		// If current link is not transaction link, it checks and retrieves transaction from context.
		if tx := gform.TXFromCtx(ctx, d.GetGroup()); tx != nil {
			link = tx
		}
	}

	// Check if it is an insert operation with primary key.
	if value := ctx.Value(internalPrimaryKeyInCtx); value != nil {
		var ok bool
		pkField, ok = value.(gform.TableField)
		if !ok {
			isUseCoreDoExec = true
		}
	} else {
		isUseCoreDoExec = true
	}

	// check if it is an insert operation.
	if !isUseCoreDoExec && pkField.Name != "" && strings.Contains(sql, "INSERT INTO") {
		primaryKey = pkField.Name
		sql += " RETURNING " + primaryKey
	} else {
		// use default DoExec
		return d.Core.DoExec(ctx, link, sql, args...)
	}

	// Only the insert operation with primary key can execute the following code

	if d.GetConfig().ExecTimeout > 0 {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(ctx, d.GetConfig().ExecTimeout)
		defer cancelFunc()
	}

	// Sql filtering.
	sql, args = d.FormatSqlBeforeExecuting(sql, args)
	sql, args, err = d.DoFilter(ctx, link, sql, args)
	if err != nil {
		return nil, err
	}

	// Link execution.
	var out gform.DoCommitOutput
	out, err = d.DoCommit(ctx, gform.DoCommitInput{
		Link:          link,
		Sql:           sql,
		Args:          args,
		Stmt:          nil,
		Type:          gform.SqlTypeQueryContext,
		IsTransaction: link.IsTransaction(),
	})

	if err != nil {
		return nil, err
	}
	affected := len(out.Records)
	if affected > 0 {
		if !strings.Contains(pkField.Type, "int") {
			return Result{
				affected:     int64(affected),
				lastInsertId: 0,
				lastInsertIdError: gerror.NewCodef(
					gcode.CodeNotSupported,
					"LastInsertId is not supported by primary key type: %s", pkField.Type),
			}, nil
		}

		if out.Records[affected-1][primaryKey] != nil {
			lastInsertId := out.Records[affected-1][primaryKey].Int64()
			return Result{
				affected:     int64(affected),
				lastInsertId: lastInsertId,
			}, nil
		}
	}

	return Result{}, nil
}
