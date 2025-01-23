/**
* 多种语言提示语言
* 添加提示语在resource/locale/xx/sysmsg.toml 中编辑
 */
package gf

import (
	"gofly/utils/tools/gfile"
	"gofly/utils/tools/gjson"
	"os"
	"path/filepath"
)

var (
	MassageDate map[string]*gjson.Json = make(map[string]*gjson.Json)
)

type GfLocale struct {
	LocaleType string // language type
}

// GetTomlData get toml file data
func GetTomlData(lange string) {
	path, _ := os.Getwd()
	install_apppath := filepath.Join(path, "resource/locale/"+lange+"/sysmsg.toml")
	jstr, _ := gjson.LoadContent(gfile.GetBytes(install_apppath))
	MassageDate[lange] = jstr
}

// Set language type
func (l *GfLocale) SetLanguage(localeType string) *GfLocale {
	if localeType == "" {
		localeType = "en-US"
	}
	l.LocaleType = localeType
	return l
}

// Get the corresponding language prompt
func (l *GfLocale) Message(msgkey string) string {
	if MassageDate[l.LocaleType] == nil {
		GetTomlData(l.LocaleType)
	}
	return MassageDate[l.LocaleType].Get(msgkey).String()
}

// Get examples of language prompts corresponding to the language
// 获取对语言对应语种提示语实例
func LocaleMsg() *GfLocale {
	l := &GfLocale{}
	l.LocaleType = "en-US"
	return l
}
