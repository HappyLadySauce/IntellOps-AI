package gins

import (
	"gofly/utils/tools/gcfg"
)

const (
	frameCoreComponentNameDatabase = "gf.core.component.database"
	frameCoreComponentNameLogger   = "gf.core.component.logger"
	frameCoreComponentNameRedis    = "gf.core.component.redis"
)

// Config returns an instance of View with default settings.
// The parameter `name` is the name for the instance.
func Config(name ...string) *gcfg.Config {
	return gcfg.Instance(name...)
}
