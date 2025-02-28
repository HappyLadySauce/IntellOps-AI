package mysql

//mysql driver
import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"gofly/utils/gform"

	"gofly/utils/tools/gcode"
	"gofly/utils/tools/gerror"

	"gofly/utils/tools/gregex"
	"gofly/utils/tools/gutil"

	"gofly/utils/gf"
)

// Driver is the driver for mysql database.
type Driver struct {
	*gform.Core
}

const (
	quoteChar = "`"
)

func init() {
	var (
		err         error
		driverObj   = New()
		driverNames = gf.SliceStr{"mysql", "mariadb", "tidb"}
	)
	for _, driverName := range driverNames {
		if err = gform.Register(driverName, driverObj); err != nil {
			panic(err)
		}
	}
}

// New create and returns a driver that implements gform.Driver, which supports operations for MySQL.
func New() gform.Driver {
	return &Driver{}
}

// New creates and returns a database object for mysql.
// It implements the interface of gform.Driver for extra database driver installation.
func (d *Driver) New(core *gform.Core, node *gform.ConfigNode) (gform.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

// Open creates and returns an underlying sql.DB object for mysql.
// Note that it converts time.Time argument to local timezone in default.
func (d *Driver) Open(config *gform.ConfigNode) (db *sql.DB, err error) {
	var (
		source               string
		underlyingDriverName = "mysql"
	)
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	if config.Link != "" {
		// ============================================================================
		// Deprecated from v2.2.0.
		// ============================================================================
		source = config.Link
		// Custom changing the schema in runtime.
		if config.Dbname != "" {
			source, _ = gregex.ReplaceString(`/([\w\.\-]+)+`, "/"+config.Dbname, source)
		}
	} else {
		// TODO: Do not set charset when charset is not specified (in v2.5.0)
		source = fmt.Sprintf(
			"%s:%s@%s(%s:%s)/%s?charset=%s",
			config.Username, config.Password, config.Protocol, config.Hostname, config.Hostport, config.Dbname, config.Charset,
		)
		if config.Timezone != "" {
			if strings.Contains(config.Timezone, "/") {
				config.Timezone = url.QueryEscape(config.Timezone)
			}
			source = fmt.Sprintf("%s&loc=%s", source, config.Timezone)
		}
		//链接成功设置数据库mode
		if config.Sqlmode != "" {
			source = fmt.Sprintf("%s&sql_mode=%s", source, config.Sqlmode)
		}
		if config.Extra != "" {
			source = fmt.Sprintf("%s&%s", source, config.Extra)
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

// GetChars returns the security char for this type of database.
func (d *Driver) GetChars() (charLeft string, charRight string) {
	return quoteChar, quoteChar
}

// DoFilter handles the sql before posts it to database.
func (d *Driver) DoFilter(ctx context.Context, link gform.Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) {
	return d.Core.DoFilter(ctx, link, sql, args)
}

// Tables retrieves and returns the tables of current schema.
// It's mainly used in cli tool chain for automatically generating the models.
func (d *Driver) Tables(ctx context.Context, schema ...string) (tables []string, err error) {
	var result gform.Result
	link, err := d.SlaveLink(schema...)
	if err != nil {
		return nil, err
	}
	result, err = d.DoSelect(ctx, link, `SHOW TABLES`)
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

// TableFields retrieves and returns the fields' information of specified table of current
// schema.
//
// The parameter `link` is optional, if given nil it automatically retrieves a raw sql connection
// as its link to proceed necessary sql query.
//
// Note that it returns a map containing the field name and its corresponding fields.
// As a map is unsorted, the TableField struct has a "Index" field marks its sequence in
// the fields.
//
// It's using cache feature to enhance the performance, which is never expired util the
// process restarts.
func (d *Driver) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*gform.TableField, err error) {
	var (
		result     gform.Result
		link       gform.Link
		usedSchema = gutil.GetOrDefaultStr(d.GetSchema(), schema...)
	)
	if link, err = d.SlaveLink(usedSchema); err != nil {
		return nil, err
	}
	result, err = d.DoSelect(
		ctx, link,
		fmt.Sprintf(`SHOW FULL COLUMNS FROM %s`, d.QuoteWord(table)),
	)
	if err != nil {
		return nil, err
	}
	fields = make(map[string]*gform.TableField)
	for i, m := range result {
		fields[m["Field"].String()] = &gform.TableField{
			Index:   i,
			Name:    m["Field"].String(),
			Type:    m["Type"].String(),
			Null:    m["Null"].Bool(),
			Key:     m["Key"].String(),
			Default: m["Default"].Val(),
			Extra:   m["Extra"].String(),
			Comment: m["Comment"].String(),
		}
	}
	return fields, nil
}
