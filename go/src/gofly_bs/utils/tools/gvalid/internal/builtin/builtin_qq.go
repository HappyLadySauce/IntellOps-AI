package builtin

import (
	"errors"

	"gofly/utils/tools/gregex"
)

// RuleQQ implements `qq` rule:
// Tencent QQ number.
//
// Format: qq
type RuleQQ struct{}

func init() {
	Register(RuleQQ{})
}

func (r RuleQQ) Name() string {
	return "qq"
}

func (r RuleQQ) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` is not a valid QQ number"
}

func (r RuleQQ) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`^[1-9][0-9]{4,}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
