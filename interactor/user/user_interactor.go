package user

import (
	"context"

	"github.com/keshvarideveloper/wastop/contract"
	"github.com/keshvarideveloper/wastop/dto"
	"github.com/keshvarideveloper/wastop/entity"
)

type Interactor struct {
	store contract.UserStore
}

func New(store contract.UserStore) Interactor {
	return Interactor{store: store}
}

func (i Interactor) SignUpUser(ctx context.Context, req dto.CreateUserRequest) (dto.CreateUserResponse, error) {

	user := entity.User{
		FullName: req.FullName,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Password: req.Password,
		Bio:      req.Bio,
		Photo:    req.Photo,
	}
	createUser, err := i.store.CreateUser(ctx, user)
	if err != nil {
		return dto.CreateUserResponse{}, err
	}
	return dto.CreateUserResponse{User: createUser}, nil
}
