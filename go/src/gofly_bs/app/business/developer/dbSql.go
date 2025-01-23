package developer

import (
	"gofly/utils/gf"
	"os"
)

// 导出数据库数据sql文件
func ExecSqlFile(tables []string, pathname string) {
	f, _ := os.Create(pathname)
	_ = gf.DBDump(
		gf.WithDropTable(),    // Option: Delete table before create (Default: Not delete table)
		gf.WithData(),         // Option: Dump Data (Default: Only dump table schema)
		gf.WithTables(tables), // Option: Dump Tables (Default: All tables)
		gf.WithWriter(f),      // Option: Writer (Default: os.Stdout)
	)
	f.Close()
}

// 导入sql文件
func ImportSqlFile(SqlPath string) error {
	return gf.ImportSql(SqlPath)
}
