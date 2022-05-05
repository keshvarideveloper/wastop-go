package dto

import "github.com/keshvarideveloper/wastop/entity"

type CreateUserRequest struct {
	FullName string `json:"fullName"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
	Bio      string `json:"bio"`
}

type CreateUserResponse struct {
	User entity.User `json:"user"`
}

type UpdateUserRequest struct {
	ID       uint   `json:"id"`
	FullName string `json:"fullName"`
	Photo    string `json:"photo"`
	Bio      string `json:"bio"`
}

type UpdateUserResponse struct {
	User entity.User `json:"user"`
}
