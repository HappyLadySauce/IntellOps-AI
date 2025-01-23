package builtin

import (
	"errors"

	"gofly/utils/tools/gipv6"
)

// RuleIpv6 implements `ipv6` rule:
// IPv6.
//
// Format: ipv6
type RuleIpv6 struct{}

func init() {
	Register(RuleIpv6{})
}

func (r RuleIpv6) Name() string {
	return "ipv6"
}

func (r RuleIpv6) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` is not a valid IPv6 address"
}

func (r RuleIpv6) Run(in RunInput) error {
	var (
		ok    bool
		value = in.Value.String()
	)
	if ok = gipv6.Validate(value); !ok {
		return errors.New(in.Message)
	}
	return nil
}
