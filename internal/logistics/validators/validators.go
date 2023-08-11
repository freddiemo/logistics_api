package validators

import (
	"github.com/go-playground/validator/v10"

	"regexp"
)

func IsVehiclePlate(field validator.FieldLevel) bool {
	match, err := regexp.MatchString("^[A-Z]{3}[0-9]{3}$", field.Field().String())

	return err == nil && match
}

func IsGuideNumber(field validator.FieldLevel) bool {
	match, err := regexp.MatchString("^[A-Za-z0-9]{10}$", field.Field().String())

	return err == nil && match
}
