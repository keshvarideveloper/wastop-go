package validator

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/keshvarideveloper/wastop/dto"
)

func ValidateCreateUser(req dto.CreateUserRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.FullName,
			validation.Required,
		),
		validation.Field(&req.Email,
			validation.Required,
			is.Email),
		validation.Field(&req.Mobile,
			is.E164),
		validation.Field(&req.Password,
			validation.Required,
		),
	)
}
