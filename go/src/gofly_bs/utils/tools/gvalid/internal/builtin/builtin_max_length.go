package builtin

import (
	"errors"
	"strconv"

	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gstr"
)

// RuleMaxLength implements `max_length` rule:
// Length is equal or lesser than :max.
// The length is calculated using unicode string, which means one chinese character or letter both has the length of 1.
//
// Format: max_length:max
type RuleMaxLength struct{}

func init() {
	Register(RuleMaxLength{})
}

func (r RuleMaxLength) Name() string {
	return "max_length"
}

func (r RuleMaxLength) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` length must be equal or lesser than {max}"
}

func (r RuleMaxLength) Run(in RunInput) error {
	var (
		valueRunes = gconv.Runes(in.Value.String())
		valueLen   = len(valueRunes)
	)
	max, err := strconv.Atoi(in.RulePattern)
	if valueLen > max || err != nil {
		return errors.New(gstr.Replace(in.Message, "{max}", strconv.Itoa(max)))
	}
	return nil
}
