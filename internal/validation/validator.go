package validation

import (
	"errors"
	"github.com/go-playground/validator/v10"
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
		return errors.New("validação falhou: " + stringJoin(validationErrors, ", "))
	}
	return nil
}

func stringJoin(arr []string, sep string) string {
	result := ""
	for i, v := range arr {
		if i == len(arr)-1 {
			result += v
		} else {
			result += v + sep
		}
	}
	return result
}
