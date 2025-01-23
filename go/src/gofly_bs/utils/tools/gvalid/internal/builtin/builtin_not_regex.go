package builtin

import (
	"errors"

	"gofly/utils/tools/gregex"
)

// RuleNotRegex implements `not_regex` rule:
// Value should not match custom regular expression pattern.
//
// Format: not_regex:pattern
type RuleNotRegex struct{}

func init() {
	Register(RuleNotRegex{})
}

func (r RuleNotRegex) Name() string {
	return "not_regex"
}

func (r RuleNotRegex) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` should not be in regex of: {pattern}"
}

func (r RuleNotRegex) Run(in RunInput) error {
	if gregex.IsMatchString(in.RulePattern, in.Value.String()) {
		return errors.New(in.Message)
	}
	return nil
}
