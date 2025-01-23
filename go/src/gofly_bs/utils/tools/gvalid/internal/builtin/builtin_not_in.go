package builtin

import (
	"errors"
	"strings"

	"gofly/utils/tools/gstr"
)

// RuleNotIn implements `not_in` rule:
// Value should not be in: value1,value2,...
//
// Format: not_in:value1,value2,...
type RuleNotIn struct{}

func init() {
	Register(RuleNotIn{})
}

func (r RuleNotIn) Name() string {
	return "not_in"
}

func (r RuleNotIn) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` must not be in range: {pattern}"
}

func (r RuleNotIn) Run(in RunInput) error {
	var (
		ok    = true
		value = in.Value.String()
	)
	for _, rulePattern := range gstr.SplitAndTrim(in.RulePattern, ",") {
		if in.Option.CaseInsensitive {
			ok = !strings.EqualFold(value, strings.TrimSpace(rulePattern))
		} else {
			ok = strings.Compare(value, strings.TrimSpace(rulePattern)) != 0
		}
		if !ok {
			return errors.New(in.Message)
		}
	}
	return nil
}
