package contract

import (
	"context"

	"github.com/keshvarideveloper/wastop/entity"
)

type UserStore interface {
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUserById(ctx context.Context, userId uint) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}
