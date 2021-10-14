package user

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"gorm.io/gorm"
)

type RepositoryUser interface {
	Create(ctx context.Context, tx *gorm.DB, user User) User
	GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) User
}

type repositoryUser struct {
}

func NewRepositoryUser() *repositoryUser {
	return &repositoryUser{}
}

func (r *repositoryUser) Create(ctx context.Context, tx *gorm.DB, user User) User {
	err := tx.WithContext(ctx).Create(&user).Error
	helper.HandleError(err)

	return user
}

func (r *repositoryUser) GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) User {
	user := User{}
	err := tx.WithContext(ctx).Where("email = ?", email).Find(&user).Error
	helper.HandleError(err)

	return user
}
