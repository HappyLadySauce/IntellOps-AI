package builtin

import (
	"errors"

	"gofly/utils/tools/gregex"
)

// RulePostcode implements `postcode` rule:
// Postcode number.
//
// Format: postcode
type RulePostcode struct{}

func init() {
	Register(RulePostcode{})
}

func (r RulePostcode) Name() string {
	return "postcode"
}

func (r RulePostcode) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` is not a valid postcode format"
}

func (r RulePostcode) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`^\d{6}$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
