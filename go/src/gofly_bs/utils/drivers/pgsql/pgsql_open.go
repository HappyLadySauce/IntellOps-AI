package pgsql

import (
	"database/sql"
	"fmt"

	"gofly/utils/gform"
	"gofly/utils/tools/gcode"
	"gofly/utils/tools/gerror"
	"gofly/utils/tools/gregex"
	"gofly/utils/tools/gstr"
)

// Open creates and returns an underlying sql.DB object for pgsql.
// https://pkg.go.dev/github.com/lib/pq
func (d *Driver) Open(config *gform.ConfigNode) (db *sql.DB, err error) {
	var (
		source               string
		underlyingDriverName = "postgres"
	)
	if config.Link != "" {
		// ============================================================================
		// Deprecated from v2.2.0.
		// ============================================================================
		source = config.Link
		// Custom changing the schema in runtime.
		if config.Dbname != "" {
			source, _ = gregex.ReplaceString(`dbname=([\w\.\-]+)+`, "dbname="+config.Dbname, source)
		}
	} else {
		if config.Dbname != "" {
			source = fmt.Sprintf(
				"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
				config.Username, config.Password, config.Hostname, config.Hostport, config.Dbname,
			)
		} else {
			source = fmt.Sprintf(
				"user=%s password=%s host=%s port=%s sslmode=disable",
				config.Username, config.Password, config.Hostname, config.Hostport,
			)
		}

		if config.Namespace != "" {
			source = fmt.Sprintf("%s search_path=%s", source, config.Namespace)
		}

		if config.Timezone != "" {
			source = fmt.Sprintf("%s timezone=%s", source, config.Timezone)
		}

		if config.Extra != "" {
			var extraMap map[string]interface{}
			if extraMap, err = gstr.Parse(config.Extra); err != nil {
				return nil, err
			}
			for k, v := range extraMap {
				source += fmt.Sprintf(` %s=%s`, k, v)
			}
		}
	}
	if db, err = sql.Open(underlyingDriverName, source); err != nil {
		err = gerror.WrapCodef(
			gcode.CodeDbOperationError, err,
			`sql.Open failed for driver "%s" by source "%s"`, underlyingDriverName, source,
		)
		return nil, err
	}
	return
}
