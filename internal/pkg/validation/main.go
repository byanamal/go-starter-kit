package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)

	for _, fe := range err.(validator.ValidationErrors) {
		field := strings.ToLower(fe.Field())
		errors[field] = validationMessage(field, fe)
	}

	return errors
}

func validationMessage(field string, err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email", field)
	case "min":
		return fmt.Sprintf("%s minimum %s characters", field, err.Param())
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}
