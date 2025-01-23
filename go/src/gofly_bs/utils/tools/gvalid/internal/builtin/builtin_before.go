package builtin

import (
	"errors"

	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gstr"
	"gofly/utils/tools/gutil"
)

// RuleBefore implements `before` rule:
// The datetime value should be after the value of field `field`.
//
// Format: before:field
type RuleBefore struct{}

func init() {
	Register(RuleBefore{})
}

func (r RuleBefore) Name() string {
	return "before"
}

func (r RuleBefore) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` must be before field {field1} value `{value1}`"
}

func (r RuleBefore) Run(in RunInput) error {
	var (
		fieldName, fieldValue = gutil.MapPossibleItemByKey(in.Data.Map(), in.RulePattern)
		valueDatetime         = in.Value.Time()
		fieldDatetime         = gconv.Time(fieldValue)
	)
	if valueDatetime.Before(fieldDatetime) {
		return nil
	}
	return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
		"{field1}": fieldName,
		"{value1}": gconv.String(fieldValue),
	}))
}
