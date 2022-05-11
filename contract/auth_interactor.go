package contract

import (
	"context"

	"github.com/keshvarideveloper/wastop/dto"
)

type AuthInteractor interface {
	SignupUser(context.Context, dto.SignupUserRequest) (dto.SignupUserResponse, error)
	LoginUserWithEmail(context.Context, dto.LoginUserWithEmailRequest) (dto.LoginUserWithEmailResponse, error)
}
