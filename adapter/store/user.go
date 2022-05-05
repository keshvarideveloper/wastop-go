package store

import (
	"context"

	"github.com/keshvarideveloper/wastop/entity"
)

func (m MySQLStore) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	u := mapFromUserEntity(user)
	if err := m.db.WithContext(ctx).Create(&u).Error; err != nil {
		return entity.User{}, err
	}
	return mapToUserEntity(u), nil
}

func (m MySQLStore) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	u := mapFromUserEntity(user)
	if err := m.db.WithContext(ctx).Create(&u).Error; err != nil {
		return entity.User{}, err
	}
	return mapToUserEntity(u), nil
}
