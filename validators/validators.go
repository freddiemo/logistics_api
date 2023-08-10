package validators

import (
	"github.com/go-playground/validator/v10"

	"net/mail"
)

func IsEmail(field validator.FieldLevel) bool {
	email := field.Field().String()
	emailAddress, err := mail.ParseAddress(email)

	return err == nil && emailAddress.Address == email
}
