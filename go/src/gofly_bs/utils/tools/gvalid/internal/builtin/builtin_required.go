package builtin

import (
	"errors"
	"reflect"

	"gofly/utils/tools/gconv"
)

// RuleRequired implements `required` rule.
// Format: required
type RuleRequired struct{}

func init() {
	Register(RuleRequired{})
}

func (r RuleRequired) Name() string {
	return "required"
}

func (r RuleRequired) Message() string {
	return MassageDate[LocaleType].Get("builtin_" + r.Name()).String()
	// return "The {field} field is required"
}

func (r RuleRequired) Run(in RunInput) error {
	if isRequiredEmpty(in.Value.Val()) {
		return errors.New(in.Message)
	}
	return nil
}

func isRequiredEmpty(value interface{}) bool {
	reflectValue := reflect.ValueOf(value)
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	switch reflectValue.Kind() {
	case reflect.String, reflect.Map, reflect.Array, reflect.Slice:
		return reflectValue.Len() == 0
	}
	return gconv.String(value) == ""
}
