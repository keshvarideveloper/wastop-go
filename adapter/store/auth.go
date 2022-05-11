package store

import (
	"context"

	"github.com/keshvarideveloper/wastop/entity"
)

func (m MySQLStore) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	user := User{}
	if err := m.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return entity.User{}, err
	}

	return mapToUserEntity(user), nil
}

func (m MySQLStore) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	u := mapFromUserEntity(user)
	if err := m.db.WithContext(ctx).Create(&u).Error; err != nil {
		return entity.User{}, err
	}
	return mapToUserEntity(u), nil
}
