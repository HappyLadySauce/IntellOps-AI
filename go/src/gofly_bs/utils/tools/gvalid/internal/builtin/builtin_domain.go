package builtin

import (
	"errors"

	"gofly/utils/tools/gregex"
)

// RuleDomain implements `domain` rule:
// Domain.
//
// Format: domain
type RuleDomain struct{}

func init() {
	Register(RuleDomain{})
}

func (r RuleDomain) Name() string {
	return "domain"
}

func (r RuleDomain) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` is not a valid domain format"
}

func (r RuleDomain) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`^([0-9a-zA-Z][0-9a-zA-Z\-]{0,62}\.)+([a-zA-Z]{0,62})$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
