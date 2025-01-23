package builtin

import (
	"errors"
	"time"

	"gofly/utils/tools/gtime"
)

// RuleDateFormat implements `date_format` rule:
// Custom date format.
//
// Format: date_format:format
type RuleDateFormat struct{}

func init() {
	Register(RuleDateFormat{})
}

func (r RuleDateFormat) Name() string {
	return "date_format"
}

func (r RuleDateFormat) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` does not match the format: {pattern}"
}

func (r RuleDateFormat) Run(in RunInput) error {
	type iTime interface {
		Date() (year int, month time.Month, day int)
		IsZero() bool
	}
	// support for time value, eg: gtime.Time/*gtime.Time, time.Time/*time.Time.
	if obj, ok := in.Value.Val().(iTime); ok {
		if obj.IsZero() {
			return errors.New(in.Message)
		}
		return nil
	}
	if _, err := gtime.StrToTimeFormat(in.Value.String(), in.RulePattern); err != nil {
		return errors.New(in.Message)
	}
	return nil
}
