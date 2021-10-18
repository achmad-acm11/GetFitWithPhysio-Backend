package user

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"gorm.io/gorm"
)

type RepositoryUser interface {
	Create(ctx context.Context, tx *gorm.DB, user User) User
	GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) User
	GetUserById(ctx context.Context, tx *gorm.DB, userId int) User
}

type repositoryUser struct {
}

func NewRepositoryUser() *repositoryUser {
	return &repositoryUser{}
}

// Query SQL Create User
func (r *repositoryUser) Create(ctx context.Context, tx *gorm.DB, user User) User {
	err := tx.WithContext(ctx).Create(&user).Error
	helper.HandleError(err)

	return user
}

// Query SQL Get User By Email
func (r *repositoryUser) GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) User {
	user := User{}
	err := tx.WithContext(ctx).Where("email = ?", email).Find(&user).Error
	helper.HandleError(err)

	return user
}

// Query SQL Get User By Id
func (r *repositoryUser) GetUserById(ctx context.Context, tx *gorm.DB, userId int) User {
	user := User{}

	err := tx.WithContext(ctx).Where("id = ?", userId).Find(&user).Error
	helper.HandleError(err)

	return user
}
