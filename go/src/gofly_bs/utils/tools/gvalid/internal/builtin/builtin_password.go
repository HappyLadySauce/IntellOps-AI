package builtin

import (
	"errors"

	"gofly/utils/tools/gregex"
)

// RulePassword implements `password` rule:
// Universal password format rule1:
// Containing any visible chars, length between 6 and 18.
//
// Format: password
type RulePassword struct{}

func init() {
	Register(RulePassword{})
}

func (r RulePassword) Name() string {
	return "password"
}

func (r RulePassword) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` is not a valid password format"
}

func (r RulePassword) Run(in RunInput) error {
	if !gregex.IsMatchString(`^[\w\S]{6,18}$`, in.Value.String()) {
		return errors.New(in.Message)
	}
	return nil
}
