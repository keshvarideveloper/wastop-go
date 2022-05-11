package user

import (
	"context"

	"github.com/keshvarideveloper/wastop/contract"
	"github.com/keshvarideveloper/wastop/dto"
)

type Interactor struct {
	store contract.UserStore
}

func New(store contract.UserStore) Interactor {
	return Interactor{
		store,
	}
}

func (i Interactor) UpdateProfile(ctx context.Context, req dto.UpdateProfileRequest) (dto.UpdateUserResponse, error) {
	user, err := i.store.GetUserById(ctx, req.ID)
	if err != nil {
		return dto.UpdateUserResponse{}, err
	}
	user.FullName = req.FullName
	user.Bio = req.Bio
	user.Photo = req.Bio

	updatedUser, err := i.store.UpdateUser(ctx, user)
	if err != nil {
		return dto.UpdateUserResponse{}, err
	}
	return dto.UpdateUserResponse{User: updatedUser}, nil
}

func (i Interactor) GetProfile(ctx context.Context, req dto.GetProfileRequest) (dto.GetProfileResponse, error) {

	user, err := i.store.GetUserById(ctx, req.ID)
	if err != nil {
		return dto.GetProfileResponse{}, err
	}
	return dto.GetProfileResponse{FullName: user.FullName, Email: user.Email,
		Mobile: user.Mobile, Photo: user.Photo, Bio: user.Bio}, nil
}
