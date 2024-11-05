package middlewares

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type validationError struct {
	HasError bool
	Field    string
	Tag      string
	Value    interface{}
}

type customValidator struct {
	validator *validator.Validate
}

var validate = validator.New()

func (v customValidator) validateAndFormatErrors(params interface{}) *map[string]string {
	if errs := v.validate(params); len(errs) > 0 && errs[0].HasError {
		errorMessages := make(map[string]string)

		for _, err := range errs {
			errorMessages[err.Field] = fmt.Sprintf("%s field failed. Validation: '%s'", err.Field, err.Tag)
		}

		return &errorMessages
	}

	return nil
}

func (v customValidator) validate(data interface{}) []validationError {
	var validationErrors []validationError

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var ve validationError

			ve.Field = err.Field()
			ve.Tag = err.Tag()
			ve.Value = err.Value()
			ve.HasError = true

			validationErrors = append(validationErrors, ve)
		}
	}

	return validationErrors
}
