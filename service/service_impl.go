package service

import (
	"GetfitWithPhysio-backend/helper"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type ServiceImpl interface {
	GetAllService(ctx context.Context) []ServiceResponse
}

type serviceImpl struct {
	repo     RepositoryService
	db       *sql.DB
	validate *validator.Validate
}

func NewServiceImpl(repo RepositoryService, db *sql.DB, validate *validator.Validate) *serviceImpl {
	return &serviceImpl{
		repo:     repo,
		db:       db,
		validate: validate,
	}
}

func (s *serviceImpl) GetAllService(ctx context.Context) []ServiceResponse {
	// Start Transaction
	tx, err := s.db.Begin()
	helper.HandleError(err)
	defer helper.CommitOrRollback(tx)

	services := s.repo.GetAll(ctx, tx)

	return MapServicesResponse(services)
}
