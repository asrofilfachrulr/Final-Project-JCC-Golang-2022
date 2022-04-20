package validation

import (
	"log"
	// error message for bahasa indonesia

	// universal translator for bahasa indonesia

	"gopkg.in/go-playground/validator.v9"
)

func RegisterValidation(v *validator.Validate, fn func() (string, validator.Func)) {
	log.Println("Registering custom validation...")
	v.RegisterValidation(fn())
}

func RegisterAll(v *validator.Validate) {
	customValidationFuncs := []func() (string, validator.Func){
		passwd,
	}

	for _, fn := range customValidationFuncs {
		RegisterValidation(v, fn)
	}

}
