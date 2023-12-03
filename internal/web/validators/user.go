package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrEmailRequired     = fiber.NewError(fiber.StatusBadRequest, "Email is required")
	ErrEmailIncorrect    = fiber.NewError(fiber.StatusBadRequest, "Incorrect email")
	ErrPasswordMissmatch = fiber.NewError(fiber.StatusBadRequest, "Passwords mismatch")
	ErrPasswordRequired  = fiber.NewError(fiber.StatusBadRequest, "Password is required")
	ErrPasswordShort     = fiber.NewError(fiber.StatusBadRequest, "Password is too short")
)

var userValidateMap = map[string]validatorFunc{
	"Email": func(e validator.FieldError) error {
		switch e.Tag() {
		case required:
			return ErrEmailRequired
		case "email":
			return ErrEmailIncorrect
		}
		return nil
	},
	"Password": func(e validator.FieldError) error {
		switch e.Tag() {
		case required:
			return ErrPasswordRequired
		case "min":
			return ErrPasswordShort
		}
		return nil
	},
	"PasswordConfirm": func(e validator.FieldError) error {
		switch e.Tag() {
		case required:
			return ErrPasswordRequired
		case "eqfield":
			return ErrPasswordMissmatch
		}
		return nil
	},
}

func (v *Validator) SwitchUserValidate(user any) error {
	if err := v.validate.Struct(user); err != nil {
		if err, ok := err.(validator.ValidationErrors); ok { //nolint:errorlint // This example from official doc
			for _, e := range err {
				if err := userValidateMap[e.Field()](e); err != nil {
					return err
				}
			}
			return err
		}
		return fiber.ErrInternalServerError
	}
	return nil
}

// func (v *Validator) SwitchUserValidate(err error) error {
// 	if err, ok := err.(validator.ValidationErrors); ok { //nolint:errorlint // This example from official doc
// 		for _, e := range err {
// 			switch e.Field() {
// 			case "Email":
// 				switch e.Tag() {
// 				case required:
// 					return ErrEmailRequired
// 				case "email":
// 					return ErrEmailIncorrect
// 				}
// 			case "Password":
// 				switch e.Tag() {
// 				case required:
// 					return ErrPasswordRequired
// 				case "min":
// 					return ErrPasswordShort
// 				}
// 			case "PasswordConfirm":
// 				switch e.Tag() {
// 				case required:
// 					return ErrPasswordRequired
// 				case "eqfield":
// 					return ErrPasswordMissmatch
// 				}
// 			}
// 		}
// 		return err
// 	}
// 	return fiber.ErrInternalServerError
// }
