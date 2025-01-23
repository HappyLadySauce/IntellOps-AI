package builtin

import (
	"errors"
	"strconv"

	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gutil"

	"gofly/utils/tools/gstr"
)

// RuleLT implements `lt` rule:
// Lesser than `field`.
// It supports both integer and float.
//
// Format: lt:field
type RuleLT struct{}

func init() {
	Register(RuleLT{})
}

func (r RuleLT) Name() string {
	return "lt"
}

func (r RuleLT) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} value `{value}` must be lesser than field {field1} value `{value1}`"
}

func (r RuleLT) Run(in RunInput) error {
	var (
		fieldName, fieldValue = gutil.MapPossibleItemByKey(in.Data.Map(), in.RulePattern)
		fieldValueN, err1     = strconv.ParseFloat(gconv.String(fieldValue), 10)
		valueN, err2          = strconv.ParseFloat(in.Value.String(), 10)
	)

	if valueN >= fieldValueN || err1 != nil || err2 != nil {
		return errors.New(gstr.ReplaceByMap(in.Message, map[string]string{
			"{field1}": fieldName,
			"{value1}": gconv.String(fieldValue),
		}))
	}
	return nil
}
