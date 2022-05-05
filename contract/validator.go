package contract

import (
	"github.com/keshvarideveloper/wastop/dto"
)

type (
	ValidateCreateUser func(req dto.CreateUserRequest) error
)
