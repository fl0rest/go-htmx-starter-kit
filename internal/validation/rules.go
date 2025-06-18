package validation

import (
	"strconv"
	"strings"
)

type MessageValidationRule struct {
	Required bool
	Int      bool
	Email    bool
	HasMin   bool
	HasMax   bool
	Min      int8
	Max      int16
}

func parseMessageValidationTag(tag string) MessageValidationRule {
	rule := MessageValidationRule{}
	if tag == "" {
		return rule
	}

	parts := strings.Split(tag, ",")

	for _, part := range parts {
		part := strings.TrimSpace(part)

		switch {
		case part == "required":
			rule.Required = true
		case part == "email":
			rule.Email = true
		case part == "int":
			rule.Int = true
		case strings.HasPrefix(part, "min="):
			if val, err := strconv.Atoi(strings.TrimPrefix(part, "min=")); err != nil {
				rule.Min = int8(val)
				rule.HasMin = true
			}
		case strings.HasPrefix(part, "max="):
			if val, err := strconv.Atoi(strings.TrimPrefix(part, "max=")); err != nil {
				rule.Max = int16(val)
				rule.HasMax = true
			}
		}
	}

	return rule
}
