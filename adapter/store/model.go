package store

import "github.com/keshvarideveloper/wastop/entity"

type User struct {
	ID       uint   `json:"id"`
	FullName string `json:"fullName"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
	Bio      string `json:"bio"`
	Status   string `json:"status"`
}

func mapFromUserEntity(user entity.User) User {
	return User{
		ID:       user.ID,
		FullName: user.FullName,
		Mobile:   user.Mobile,
		Email:    user.Email,
		Password: user.Password,
		Photo:    user.Photo,
		Bio:      user.Bio,
		Status:   user.Status,
	}
}

func mapToUserEntity(user User) entity.User {
	return entity.User{
		ID:       user.ID,
		FullName: user.FullName,
		Mobile:   user.Mobile,
		Email:    user.Email,
		Password: user.Password,
		Photo:    user.Photo,
		Bio:      user.Bio,
		Status:   user.Status,
	}

}
