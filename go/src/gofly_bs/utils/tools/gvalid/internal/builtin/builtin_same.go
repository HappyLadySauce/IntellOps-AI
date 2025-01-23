package builtin

import (
	"errors"
	"strings"

	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gstr"
	"gofly/utils/tools/gutil"
)

// RuleSame implements `same` rule:
// Value should be the same as value of field.
//
// Format: same:field
type RuleSame struct{}

func init() {
	Register(RuleSame{})
}

func (r RuleSame) Name() string {
	return "same"
}

func (r RuleSame) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` must be the same as field {field1} value `{value1}`"
}

func (r RuleSame) Run(in RunInput) error {
	var (
		ok    bool
		value = in.Value.String()
	)
	fieldName, fieldValue := gutil.MapPossibleItemByKey(in.Data.Map(), in.RulePattern)
	if fieldValue != nil {
		if in.Option.CaseInsensitive {
			ok = strings.EqualFold(value, gconv.String(fieldValue))
		} else {
			ok = strings.Compare(value, gconv.String(fieldValue)) == 0
		}
	}
	if !ok {
		return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
			"{field1}": fieldName,
			"{value1}": gconv.String(fieldValue),
		}))
	}
	return nil
}
