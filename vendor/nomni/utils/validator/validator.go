package validator

import (
	"nomni/utils/api"

	validator "gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err == nil {
		return err
	}
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, err := range errs {
			return api.InvalidParamError(err.Field(), err.Tag(), errs)
		}
	}
	return err
}
func New() *Validator {
	return &Validator{validator: validator.New()}
}
