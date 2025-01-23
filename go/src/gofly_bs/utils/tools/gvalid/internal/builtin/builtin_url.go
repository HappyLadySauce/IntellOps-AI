package builtin

import (
	"errors"

	"gofly/utils/tools/gregex"
)

// RuleUrl implements `url` rule:
// URL.
//
// Format: url
type RuleUrl struct{}

func init() {
	Register(RuleUrl{})
}

func (r RuleUrl) Name() string {
	return "url"
}

func (r RuleUrl) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` is not a valid URL address"
}

func (r RuleUrl) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
