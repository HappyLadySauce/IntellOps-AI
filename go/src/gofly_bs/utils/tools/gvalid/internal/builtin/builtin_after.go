package builtin

import (
	"errors"

	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gutil"

	"gofly/utils/tools/gstr"
)

// RuleAfter implements `after` rule:
// The datetime value should be after the value of field `field`.
//
// Format: after:field
type RuleAfter struct{}

func init() {
	Register(RuleAfter{})
}

func (r RuleAfter) Name() string {
	return "after"
}

func (r RuleAfter) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` must be after field {field1} value `{value1}`"
}

func (r RuleAfter) Run(in RunInput) error {
	var (
		fieldName, fieldValue = gutil.MapPossibleItemByKey(in.Data.Map(), in.RulePattern)
		valueDatetime         = in.Value.Time()
		fieldDatetime         = gconv.Time(fieldValue)
	)
	if valueDatetime.After(fieldDatetime) {
		return nil
	}
	return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
		"{field1}": fieldName,
		"{value1}": gconv.String(fieldValue),
	}))
}
