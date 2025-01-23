package gvalid

import (
	"context"

	"gofly/utils/tools/gvalid/internal/builtin"
)

// getErrorMessageByRule retrieves and returns the error message for specified rule.
// It firstly retrieves the message from custom message map, and then checks i18n manager,
// it returns the default error message if it's not found in neither custom message map nor i18n manager.
func (v *Validator) getErrorMessageByRule(ctx context.Context, ruleKey string, customMsgMap map[string]string) string {
	content := customMsgMap[ruleKey]
	if content != "" {
		return content
	}

	// Retrieve default message according to certain rule.
	// Builtin rule message.
	if content == "" {
		if builtinRule := builtin.GetRule(ruleKey, v.localeType); builtinRule != nil {
			content = builtinRule.Message()
		}
	}
	// If there's no configured rule message, it uses default one.
	if content == "" {
		content = defaultErrorMessages[internalDefaultRuleName]
	}
	return content
}
