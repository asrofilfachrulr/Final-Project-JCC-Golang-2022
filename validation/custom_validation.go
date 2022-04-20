package validation

import "gopkg.in/go-playground/validator.v9"

func passwd() (string, validator.Func) {
	return "passwd", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) > 4
	}
}
