package contract

import (
	"context"

	"github.com/keshvarideveloper/wastop/dto"
)

type UserInteractor interface {
	UpdateProfile(context.Context, dto.UpdateProfileRequest) (dto.UpdateUserResponse, error)
	GetProfile(context.Context, dto.GetProfileRequest) (dto.GetProfileResponse, error)
}
