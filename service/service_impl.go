package service

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ServiceImpl interface {
	GetAllService(ctx context.Context) []ServiceResponse
}

type serviceImpl struct {
	repo     RepositoryService
	db       *gorm.DB
	validate *validator.Validate
}

func NewServiceImpl(repo RepositoryService, db *gorm.DB, validate *validator.Validate) *serviceImpl {
	return &serviceImpl{
		repo:     repo,
		db:       db,
		validate: validate,
	}
}

func (s *serviceImpl) GetAllService(ctx context.Context) []ServiceResponse {
	// Start Transaction
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	services := s.repo.GetAll(ctx, tx)

	return MapServicesResponse(services)
}
