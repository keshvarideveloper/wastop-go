package contract

import (
	"context"

	"github.com/keshvarideveloper/wastop/dto"
)

type UserInteractor interface {
	SignUpUser(context.Context, dto.CreateUserRequest) (dto.CreateUserResponse, error)
	UpdateUser(context.Context, dto.UpdateUserRequest) (dto.UpdateUserResponse, error)
}
