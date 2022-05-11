package dto

import "github.com/keshvarideveloper/wastop/entity"

type LoginUserWithEmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserWithEmailResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
type SignupUserRequest struct {
	FullName string `json:"fullName"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
	Bio      string `json:"bio"`
}

type SignupUserResponse struct {
	User entity.User `json:"user"`
}
