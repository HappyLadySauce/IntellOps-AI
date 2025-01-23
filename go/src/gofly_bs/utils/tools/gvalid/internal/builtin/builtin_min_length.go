package builtin

import (
	"errors"
	"strconv"

	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gstr"
)

// RuleMinLength implements `min_length` rule:
// Length is equal or greater than :min.
// The length is calculated using unicode string, which means one chinese character or letter both has the length of 1.
//
// Format: min_length:min
type RuleMinLength struct{}

func init() {
	Register(RuleMinLength{})
}

func (r RuleMinLength) Name() string {
	return "min_length"
}

func (r RuleMinLength) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` length must be equal or greater than {min}"
}

func (r RuleMinLength) Run(in RunInput) error {
	var (
		valueRunes = gconv.Runes(in.Value.String())
		valueLen   = len(valueRunes)
	)
	min, err := strconv.Atoi(in.RulePattern)
	if valueLen < min || err != nil {
		return errors.New(gstr.Replace(in.Message, "{min}", strconv.Itoa(min)))
	}
	return nil
}
