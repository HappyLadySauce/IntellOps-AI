package builtin

import (
	"os"
	"path/filepath"
	"reflect"

	"gofly/utils/tools/gfile"
	"gofly/utils/tools/gjson"
	"gofly/utils/tools/gvar"
)

type Rule interface {
	// Name returns the builtin name of the rule.
	Name() string

	// Message returns the default error message of the rule.
	Message() string

	// Run starts running the rule, it returns nil if successful, or else an error.
	Run(in RunInput) error
}

type RunInput struct {
	RuleKey     string       // RuleKey is like the "max" in rule "max: 6"
	RulePattern string       // RulePattern is like "6" in rule:"max:6"
	Field       string       // The field name of Value.
	ValueType   reflect.Type // ValueType specifies the type of the value, which might be nil.
	Value       *gvar.Var    // Value specifies the value for this rule to validate.
	Data        *gvar.Var    // Data specifies the `data` which is passed to the Validator.
	Message     string       // Message specifies the custom error message or configured i18n message for this rule.
	Option      RunOption    // Option provides extra configuration for validation rule.
}

type RunOption struct {
	CaseInsensitive bool // CaseInsensitive indicates that it does Case-Insensitive comparison in string.
}

var (
	// ruleMap stores all builtin validation rules.
	ruleMap                            = map[string]Rule{}
	LocaleType                         = "en-US"
	MassageDate map[string]*gjson.Json = make(map[string]*gjson.Json)
)

// Register registers builtin rule into manager.
func Register(rule Rule) {
	ruleMap[rule.Name()] = rule
}

// GetTomlData get toml file data
func GetTomlData(lange string) {
	path, _ := os.Getwd()
	install_apppath := filepath.Join(path, "resource/locale/"+lange+"/validation.toml")
	jstr, _ := gjson.LoadContent(gfile.GetBytes(install_apppath))
	MassageDate[lange] = jstr
}

// GetRule retrieves and returns rule by `name`.
func GetRule(name, localeType string) Rule {
	LocaleType = localeType
	if MassageDate[localeType] == nil {
		GetTomlData(localeType)
	}
	return ruleMap[name]
}
