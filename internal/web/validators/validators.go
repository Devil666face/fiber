package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
)

const (
	required = "required"
)

type validatorFunc func(e validator.FieldError) error

type Validator struct {
	validate *validator.Validate
	policy   *bluemonday.Policy
}

func New() *Validator {
	return &Validator{
		validate: validator.New(validator.WithRequiredStructEnabled()),
		policy:   bluemonday.StrictPolicy(),
	}
}

func (v *Validator) Validate() *validator.Validate {
	return v.validate
}

func (v *Validator) validateStringField(field string) bool {
	return len([]rune(field)) == len([]rune(v.policy.Sanitize(field)))
}

func (v *Validator) ValidateInputs(fields ...string) bool {
	for _, f := range fields {
		if !v.validateStringField(f) {
			return false
		}
	}
	return true
}
