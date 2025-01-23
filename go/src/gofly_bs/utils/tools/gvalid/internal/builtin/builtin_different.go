package builtin

import (
	"errors"
	"strings"

	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gstr"
	"gofly/utils/tools/gutil"
)

// RuleDifferent implements `different` rule:
// Value should be different from value of field.
//
// Format: different:field
type RuleDifferent struct{}

func init() {
	Register(RuleDifferent{})
}

func (r RuleDifferent) Name() string {
	return "different"
}

func (r RuleDifferent) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` must be different from field {field1} value `{value1}`"
}

func (r RuleDifferent) Run(in RunInput) error {
	var (
		ok    = true
		value = in.Value.String()
	)
	fieldName, fieldValue := gutil.MapPossibleItemByKey(in.Data.Map(), in.RulePattern)
	if fieldValue != nil {
		if in.Option.CaseInsensitive {
			ok = !strings.EqualFold(value, gconv.String(fieldValue))
		} else {
			ok = strings.Compare(value, gconv.String(fieldValue)) != 0
		}
	}
	if ok {
		return nil
	}
	return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
		"{field1}": fieldName,
		"{value1}": gconv.String(fieldValue),
	}))
}
