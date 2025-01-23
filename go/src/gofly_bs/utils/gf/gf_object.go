package gf

import (
	"gofly/utils/gform"
	"gofly/utils/tools/gcfg"
	"gofly/utils/tools/gins"
	"gofly/utils/tools/glog"
	"gofly/utils/tools/gredis"
	"gofly/utils/tools/gvalid"
	"gofly/utils/tools/gvar"
)

type (
	OrmResult = gform.Result //Go Result
	OrmRecord = gform.Record //Go Record
)

// DB returns an instance of database ORM object with specified configuration group name.
func DB(name ...string) gform.DB {
	return gins.Database(name...)
}

// Model creates and returns a model based on configuration of default database group.
func Model(tableNameOrStruct ...interface{}) *gform.Model {
	return DB("default").Model(tableNameOrStruct...)
}

// ModelRaw creates and returns a model based on a raw sql not a table.
func ModelRaw(rawSql string, args ...interface{}) *gform.Model {
	return DB().Raw(rawSql, args...)
}

// Config returns an instance of config object with specified name.
func Config(name ...string) *gcfg.Config {
	return gins.Config(name...)
}

// Cfg is alias of Config.
// See Config.
func Cfg(name ...string) *gcfg.Config {
	return Config(name...)
}

// Redis returns an instance of redis client with specified configuration group name.
func Redis(name ...string) *gredis.Redis {
	return gins.Redis(name...)
}

// Log returns an instance of glog.Logger.
// The parameter `name` is the name for the instance.
func Log(name ...string) *glog.Logger {
	return gins.Log(name...)
}

// Validator is a convenience function, which creates and returns a new validation manager object.
func Validator() *gvalid.Validator {
	return gvalid.New()
}

// 把任何类型转var
func VarNew(val interface{}) *gvar.Var {
	return gvar.New(val)
}
