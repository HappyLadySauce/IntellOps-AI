package builtin

import (
	"errors"
	"strings"

	"gofly/utils/tools/empty"
	"gofly/utils/tools/gutil"
)

// RuleRequiredWithAll implements `required-with-all` rule:
// Required if all given fields are not empty.
//
// Format:  required-with-all:field1,field2,...
// Example: required-with-all:id,name
type RuleRequiredWithAll struct{}

func init() {
	Register(RuleRequiredWithAll{})
}

func (r RuleRequiredWithAll) Name() string {
	return "required_with_all"
}

func (r RuleRequiredWithAll) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} field is required"
}

func (r RuleRequiredWithAll) Run(in RunInput) error {
	var (
		required   = true
		array      = strings.Split(in.RulePattern, ",")
		foundValue interface{}
	)
	for i := 0; i < len(array); i++ {
		_, foundValue = gutil.MapPossibleItemByKey(in.Data.Map(), array[i])
		if empty.IsEmpty(foundValue) {
			required = false
			break
		}
	}

	if required && isRequiredEmpty(in.Value.Val()) {
		return errors.New(in.Message)
	}
	return nil
}
