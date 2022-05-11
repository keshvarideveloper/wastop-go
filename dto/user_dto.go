package dto

import "github.com/keshvarideveloper/wastop/entity"

type UpdateProfileRequest struct {
	ID       uint   `json:"id"`
	FullName string `json:"fullName"`
	Photo    string `json:"photo"`
	Bio      string `json:"bio"`
}

type UpdateUserResponse struct {
	User entity.User `json:"user"`
}

type GetProfileRequest struct {
	ID uint `json:"id"`
}

type GetProfileResponse struct {
	FullName string `json:"fullName"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Photo    string `json:"photo"`
	Bio      string `json:"bio"`
}
