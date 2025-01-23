package gins

import (
	"context"
	"fmt"

	"gofly/utils/gform"

	"gofly/utils/tools/gcode"
	"gofly/utils/tools/gerror"
	"gofly/utils/tools/glog"

	"gofly/utils/tools/consts"

	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gutil"
	"gofly/utils/tools/instance"
	"gofly/utils/tools/intlog"

	"gofly/utils/tools/gcfg"
)

// Database returns an instance of database ORM object with specified configuration group name.
// Note that it panics if any error occurs duration instance creating.
func Database(name ...string) gform.DB {
	var (
		ctx   = context.Background()
		group = gform.DefaultGroupName
	)

	if len(name) > 0 && name[0] != "" {
		group = name[0]
	}
	instanceKey := fmt.Sprintf("%s.%s", frameCoreComponentNameDatabase, group)
	db := instance.GetOrSetFuncLock(instanceKey, func() interface{} {
		// It ignores returned error to avoid file no found error while it's not necessary.
		var (
			configMap     map[string]interface{}
			configNodeKey = consts.ConfigNodeNameDatabase
		)
		// It firstly searches the configuration of the instance name.
		if configData, _ := Config().Data(ctx); len(configData) > 0 {
			if v, _ := gutil.MapPossibleItemByKey(configData, consts.ConfigNodeNameDatabase); v != "" {
				configNodeKey = v
			}
		}
		if v, _ := Config().Get(ctx, configNodeKey); !v.IsEmpty() {
			configMap = v.Map()
		}
		// No configuration found, it formats and panics error.
		if len(configMap) == 0 && !gform.IsConfigured() {
			// File configuration object checks.
			var err error
			if fileConfig, ok := Config().GetAdapter().(*gcfg.AdapterFile); ok {
				if _, err = fileConfig.GetFilePath(); err != nil {
					panic(gerror.WrapCode(gcode.CodeMissingConfiguration, err,
						`configuration not found, did you miss the configuration file or misspell the configuration file name`,
					))
				}
			}
			if v, _ := Config().Get(ctx, configNodeKey); !v.IsEmpty() {
				configMap = v.Map()
			}
			// Panic if nothing found in Config object or in gform configuration.
			if len(configMap) == 0 && !gform.IsConfigured() {
				panic(gerror.NewCodef(
					gcode.CodeMissingConfiguration,
					`database initialization failed: configuration missing for database node "%s"`,
					consts.ConfigNodeNameDatabase,
				))
			}
		}

		if len(configMap) == 0 {
			configMap = make(map[string]interface{})
		}
		// Parse `m` as map-slice and adds it to global configurations for package gform.
		for g, groupConfig := range configMap {
			cg := gform.ConfigGroup{}
			switch value := groupConfig.(type) {
			case []interface{}:
				for _, v := range value {
					if node := parseDBConfigNode(v); node != nil {
						cg = append(cg, *node)
					}
				}
			case map[string]interface{}:
				if node := parseDBConfigNode(value); node != nil {
					cg = append(cg, *node)
				}
			}
			if len(cg) > 0 {
				if gform.GetConfig(group) == nil {
					intlog.Printf(ctx, "add configuration for group: %s, %#v", g, cg)
					gform.SetConfigGroup(g, cg)
				} else {
					intlog.Printf(ctx, "ignore configuration as it already exists for group: %s, %#v", g, cg)
					intlog.Printf(ctx, "%s, %#v", g, cg)
				}
			}
		}
		// Parse `m` as a single node configuration,
		// which is the default group configuration.
		if node := parseDBConfigNode(configMap); node != nil {
			cg := gform.ConfigGroup{}
			if node.Link != "" || node.Hostname != "" {
				cg = append(cg, *node)
			}
			if len(cg) > 0 {
				if gform.GetConfig(group) == nil {
					intlog.Printf(ctx, "add configuration for group: %s, %#v", gform.DefaultGroupName, cg)
					gform.SetConfigGroup(gform.DefaultGroupName, cg)
				} else {
					intlog.Printf(
						ctx,
						"ignore configuration as it already exists for group: %s, %#v",
						gform.DefaultGroupName, cg,
					)
					intlog.Printf(ctx, "%s, %#v", gform.DefaultGroupName, cg)
				}
			}
		}
		// Create a new ORM object with given configurations.
		if db, err := gform.NewByGroup(name...); err == nil {
			// Initialize logger for ORM.
			var (
				loggerConfigMap map[string]interface{}
				loggerNodeName  = fmt.Sprintf("%s.%s", configNodeKey, consts.ConfigNodeNameLogger)
			)
			if v, _ := Config().Get(ctx, loggerNodeName); !v.IsEmpty() {
				loggerConfigMap = v.Map()
			}
			if len(loggerConfigMap) == 0 {
				if v, _ := Config().Get(ctx, configNodeKey); !v.IsEmpty() {
					loggerConfigMap = v.Map()
				}
			}
			if len(loggerConfigMap) > 0 {
				if logger, ok := db.GetLogger().(*glog.Logger); ok {
					if err = logger.SetConfigWithMap(loggerConfigMap); err != nil {
						panic(err)
					}
				}
			}
			return db
		} else {
			// If panics, often because it does not find its configuration for given group.
			panic(err)
		}
	})
	if db != nil {
		return db.(gform.DB)
	}
	return nil
}

func parseDBConfigNode(value interface{}) *gform.ConfigNode {
	nodeMap, ok := value.(map[string]interface{})
	if !ok {
		return nil
	}
	var (
		node = &gform.ConfigNode{}
		err  = gconv.Struct(nodeMap, node)
	)
	if err != nil {
		panic(err)
	}
	// Find possible `Link` configuration content.
	if _, v := gutil.MapPossibleItemByKey(nodeMap, "Link"); v != nil {
		node.Link = gconv.String(v)
	}
	return node
}
