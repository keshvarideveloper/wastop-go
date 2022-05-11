package auth

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/keshvarideveloper/wastop/contract"
	"github.com/keshvarideveloper/wastop/dto"
	"github.com/keshvarideveloper/wastop/entity"
	"golang.org/x/crypto/bcrypt"
)

type customClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type Interactor struct {
	store contract.AuthStore
}

func New(store contract.AuthStore) Interactor {
	return Interactor{
		store,
	}
}
func (i Interactor) SignupUser(ctx context.Context, req dto.SignupUserRequest) (dto.SignupUserResponse, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 13)
	user := entity.User{
		FullName: req.FullName,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Password: string(bytes),
		Photo:    req.Photo,
		Bio:      req.Bio,
		Status:   entity.Registered,
	}

	createdUser, err := i.store.CreateUser(ctx, user)
	if err != nil {
		return dto.SignupUserResponse{}, err
	}
	return dto.SignupUserResponse{User: createdUser}, nil
}

func (i Interactor) LoginUserWithEmail(ctx context.Context, req dto.LoginUserWithEmailRequest) (dto.LoginUserWithEmailResponse, error) {

	user, err := i.store.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return dto.LoginUserWithEmailResponse{}, err
	}

	compareFaild := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if compareFaild != nil {
		return dto.LoginUserWithEmailResponse{}, err
	}

	claims := customClaims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Id:        strconv.FormatUint(uint64(user.ID), 10),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := accessToken.SignedString([]byte("ssdecdasdasuraseSsecreTeassxdasdt"))
	if err != nil {
		return dto.LoginUserWithEmailResponse{}, err
	}
	fmt.Println(signedToken)
	return dto.LoginUserWithEmailResponse{AccessToken: signedToken}, nil

}
