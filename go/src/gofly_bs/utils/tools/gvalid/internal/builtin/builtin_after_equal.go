package builtin

import (
	"errors"

	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gutil"

	"gofly/utils/tools/gstr"
)

// RuleAfterEqual implements `after_equal` rule:
// The datetime value should be after or equal to the value of field `field`.
//
// Format: after_equal:field
type RuleAfterEqual struct{}

func init() {
	Register(RuleAfterEqual{})
}

func (r RuleAfterEqual) Name() string {
	return "after_equal"
}

func (r RuleAfterEqual) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` must be after or equal to field {field1} value `{value1}`"
}

func (r RuleAfterEqual) Run(in RunInput) error {
	var (
		fieldName, fieldValue = gutil.MapPossibleItemByKey(in.Data.Map(), in.RulePattern)
		valueDatetime         = in.Value.Time()
		fieldDatetime         = gconv.Time(fieldValue)
	)
	if valueDatetime.After(fieldDatetime) || valueDatetime.Equal(fieldDatetime) {
		return nil
	}
	return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
		"{field1}": fieldName,
		"{value1}": gconv.String(fieldValue),
	}))
}
