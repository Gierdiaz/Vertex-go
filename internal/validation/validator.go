package validation

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(entity interface{}) error {
	err := validate.Struct(entity)
	if err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		return errors.New("validação falhou: " + strings.Join(validationErrors, ", "))
	}
	return nil
}
