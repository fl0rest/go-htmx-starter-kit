package validation

import (
	"fmt"
	"project/internal/models"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func ValidateMessage(m models.Message) []string {
	var errors []string
	v := reflect.ValueOf(m)
	t := reflect.TypeOf(m)

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	for i := range v.NumField() {
		field := v.Field(i)
		fieldType := t.Field(i)
		tag := fieldType.Tag.Get("validation")

		rule := parseMessageValidationTag(tag)
		fieldName := fieldType.Name
		fieldValue := field.String()

		if tag == "- " {
			continue
		}

		if rule.Required && strings.TrimSpace(fieldValue) == "" {
			errors = append(errors, fmt.Sprintf("%s is required", fieldName))
			continue
		}

		if !rule.Required && strings.TrimSpace(fieldValue) == "" {
			continue
		}

		if rule.Int {
			if _, err := strconv.Atoi(fieldValue); err != nil {
				errors = append(errors, fmt.Sprintf("%v is not a number", fieldValue))
				return errors
			}
		}

		if rule.Email && !emailRegex.MatchString(fieldValue) {
			errors = append(errors, fmt.Sprintf("%s is not a valid email", fieldName))
		}

		if rule.HasMin && len(fieldValue) < int(rule.Min) {
			errors = append(errors, fmt.Sprintf("%s must be at least %d characters", fieldName, rule.Min))
		}

		if rule.HasMax && len(fieldValue) > int(rule.Max) {
			errors = append(errors, fmt.Sprintf("%s must be at less than %d characters", fieldName, rule.Max))
		}
	}

	return errors
}
