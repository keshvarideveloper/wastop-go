package contract

import (
	"context"

	"github.com/keshvarideveloper/wastop/entity"
)

type AuthStore interface {
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}
