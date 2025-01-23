package builtin

import (
	"errors"

	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gutil"

	"gofly/utils/tools/gstr"
)

// RuleBeforeEqual implements `before_equal` rule:
// The datetime value should be after or equal to the value of field `field`.
//
// Format: before_equal:field
type RuleBeforeEqual struct{}

func init() {
	Register(RuleBeforeEqual{})
}

func (r RuleBeforeEqual) Name() string {
	return "before_equal"
}

func (r RuleBeforeEqual) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` must be before or equal to field {pattern}"
}

func (r RuleBeforeEqual) Run(in RunInput) error {
	var (
		fieldName, fieldValue = gutil.MapPossibleItemByKey(in.Data.Map(), in.RulePattern)
		valueDatetime         = in.Value.Time()
		fieldDatetime         = gconv.Time(fieldValue)
	)
	if valueDatetime.Before(fieldDatetime) || valueDatetime.Equal(fieldDatetime) {
		return nil
	}
	return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
		"{field1}": fieldName,
		"{value1}": gconv.String(fieldValue),
	}))
}
