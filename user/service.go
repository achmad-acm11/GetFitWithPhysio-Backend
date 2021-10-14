package user

import (
	"GetfitWithPhysio-backend/exception"
	"GetfitWithPhysio-backend/helper"
	"context"
	"errors"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ServiceUser interface {
	Login(ctx context.Context, req LoginRequest)
}

type serviceUser struct {
	repo     RepositoryUser
	db       *gorm.DB
	validate *validator.Validate
}

func NewServiceUser(repo RepositoryUser, db *gorm.DB, validate *validator.Validate) *serviceUser {
	return &serviceUser{
		repo:     repo,
		db:       db,
		validate: validate,
	}
}

func (s *serviceUser) Login(ctx context.Context, req LoginRequest) {
	// Validate
	err := s.validate.Struct(req)
	helper.HandleError(err)

	// Transaction Start
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	user := s.repo.GetUserByEmail(ctx, tx, req.Email)

	if user.Id == 0 {
		panic(exception.NewFailedLogin(errors.New("not found email").Error()))
	}

	// Check Password
	if user.Password != req.Password {
		panic(exception.NewFailedLogin(errors.New("wrong password").Error()))
	}
}
