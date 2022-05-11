package contract

import (
	"context"

	"github.com/keshvarideveloper/wastop/dto"
)

type (
	ValidateSignupUser    func(req dto.SignupUserRequest) error
	ValidateUpdateProfile func(ctx context.Context, req dto.UpdateProfileRequest) error
	ValidateGetProfile    func(ctx context.Context, req dto.GetProfileRequest) error
)
