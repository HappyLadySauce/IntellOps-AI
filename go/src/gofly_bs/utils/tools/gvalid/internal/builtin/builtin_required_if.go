package builtin

import (
	"errors"
	"strings"

	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gutil"
)

// RuleRequiredIf implements `required_if` rule:
// Required unless all given field and its value are equal.
//
// Format:  required_if:field,value,...
// Example: required_if: id,1,age,18
type RuleRequiredIf struct{}

func init() {
	Register(RuleRequiredIf{})
}

func (r RuleRequiredIf) Name() string {
	return "required_if"
}

func (r RuleRequiredIf) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} field is required"
}

func (r RuleRequiredIf) Run(in RunInput) error {
	var (
		required   = false
		array      = strings.Split(in.RulePattern, ",")
		foundValue interface{}
	)
	// It supports multiple field and value pairs.
	if len(array)%2 == 0 {
		for i := 0; i < len(array); {
			tk := array[i]
			tv := array[i+1]
			_, foundValue = gutil.MapPossibleItemByKey(in.Data.Map(), tk)
			if in.Option.CaseInsensitive {
				required = strings.EqualFold(tv, gconv.String(foundValue))
			} else {
				required = strings.Compare(tv, gconv.String(foundValue)) == 0
			}
			if required {
				break
			}
			i += 2
		}
	}

	if required && isRequiredEmpty(in.Value.Val()) {
		return errors.New(in.Message)
	}
	return nil
}
