package utils

import "github.com/go-playground/validator/v10"

type ResponseValidation struct {
	ValidationError interface{}
}

func ValidationResponse(err error) ResponseValidation {
	var res ResponseValidation
	output := map[string]interface{}{}

	for _, err := range err.(validator.ValidationErrors) {
		output[err.Field()] = err.Error()
	}

	res.ValidationError = output
	return res
}
