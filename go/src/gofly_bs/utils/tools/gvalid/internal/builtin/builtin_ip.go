package builtin

import (
	"errors"

	"gofly/utils/tools/gipv4"
	"gofly/utils/tools/gipv6"
)

// RuleIp implements `ip` rule:
// IPv4/IPv6.
//
// Format: ip
type RuleIp struct{}

func init() {
	Register(RuleIp{})
}

func (r RuleIp) Name() string {
	return "ip"
}

func (r RuleIp) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` is not a valid IP address"
}

func (r RuleIp) Run(in RunInput) error {
	var (
		ok    bool
		value = in.Value.String()
	)
	if ok = gipv4.Validate(value) || gipv6.Validate(value); !ok {
		return errors.New(in.Message)
	}
	return nil
}
