package builtin

import (
	"errors"
	"strconv"

	"gofly/utils/tools/gstr"
)

// RuleMin implements `min` rule:
// Equal or greater than :min. It supports both integer and float.
//
// Format: min:min
type RuleMin struct{}

func init() {
	Register(RuleMin{})
}

func (r RuleMin) Name() string {
	return "min"
}

func (r RuleMin) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` must be equal or greater than {min}"
}

func (r RuleMin) Run(in RunInput) error {
	var (
		min, err1    = strconv.ParseFloat(in.RulePattern, 10)
		valueN, err2 = strconv.ParseFloat(in.Value.String(), 10)
	)
	if valueN < min || err1 != nil || err2 != nil {
		return errors.New(gstr.Replace(in.Message, "{min}", strconv.FormatFloat(min, 'f', -1, 64)))
	}
	return nil
}
