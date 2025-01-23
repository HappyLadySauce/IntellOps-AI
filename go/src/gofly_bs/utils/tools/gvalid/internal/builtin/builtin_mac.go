package builtin

import (
	"errors"

	"gofly/utils/tools/gregex"
)

// RuleMac implements `mac` rule:
// MAC.
//
// Format: mac
type RuleMac struct{}

func init() {
	Register(RuleMac{})
}

func (r RuleMac) Name() string {
	return "mac"
}

func (r RuleMac) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` is not a valid MAC address"
}

func (r RuleMac) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`^([0-9A-Fa-f]{2}[\-:]){5}[0-9A-Fa-f]{2}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
