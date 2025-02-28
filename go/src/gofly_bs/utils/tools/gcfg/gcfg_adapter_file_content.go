package gcfg

import (
	"context"
	"gofly/utils/tools/intlog"
)

// SetContent sets customized configuration content for specified `file`.
// The `file` is unnecessary param, default is DefaultConfigFile.
func (a *AdapterFile) SetContent(content string, file ...string) {
	name := DefaultConfigFileName
	if len(file) > 0 {
		name = file[0]
	}
	// Clear file cache for instances which cached `name`.
	localInstances.LockFunc(func(m map[string]interface{}) {
		if customConfigContentMap.Contains(name) {
			for _, v := range m {
				if configInstance, ok := v.(*Config); ok {
					if fileConfig, ok := configInstance.GetAdapter().(*AdapterFile); ok {
						fileConfig.jsonMap.Remove(name)
					}
				}
			}
		}
		customConfigContentMap.Set(name, content)
	})
}

// GetContent returns customized configuration content for specified `file`.
// The `file` is unnecessary param, default is DefaultConfigFile.
func (a *AdapterFile) GetContent(file ...string) string {
	name := DefaultConfigFileName
	if len(file) > 0 {
		name = file[0]
	}
	return customConfigContentMap.Get(name)
}

// RemoveContent removes the global configuration with specified `file`.
// If `name` is not passed, it removes configuration of the default group name.
func (a *AdapterFile) RemoveContent(file ...string) {
	name := DefaultConfigFileName
	if len(file) > 0 {
		name = file[0]
	}
	// Clear file cache for instances which cached `name`.
	localInstances.LockFunc(func(m map[string]interface{}) {
		if customConfigContentMap.Contains(name) {
			for _, v := range m {
				if configInstance, ok := v.(*Config); ok {
					if fileConfig, ok := configInstance.GetAdapter().(*AdapterFile); ok {
						fileConfig.jsonMap.Remove(name)
					}
				}
			}
			customConfigContentMap.Remove(name)
		}
	})

	intlog.Printf(context.TODO(), `RemoveContent: %s`, name)
}

// ClearContent removes all global configuration contents.
func (a *AdapterFile) ClearContent() {
	customConfigContentMap.Clear()
	// Clear cache for all instances.
	localInstances.LockFunc(func(m map[string]interface{}) {
		for _, v := range m {
			if configInstance, ok := v.(*Config); ok {
				if fileConfig, ok := configInstance.GetAdapter().(*AdapterFile); ok {
					fileConfig.jsonMap.Clear()
				}
			}
		}
	})
	intlog.Print(context.TODO(), `RemoveConfig`)
}
