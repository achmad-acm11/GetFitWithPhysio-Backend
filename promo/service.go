package promo

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ServicePromo interface {
	GetAllService(ctx context.Context) []PromoResponse
}

type servicePromo struct {
	repo     RepositoryPromo
	db       *gorm.DB
	validate *validator.Validate
}

func NewServicePromo(repo RepositoryPromo, db *gorm.DB, validate *validator.Validate) *servicePromo {
	return &servicePromo{
		repo:     repo,
		db:       db,
		validate: validate,
	}
}

func (s *servicePromo) GetAllService(ctx context.Context) []PromoResponse {
	// Start Transaction
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	promos := s.repo.GetAll(ctx, tx)

	return MapPromosResponse(promos)
}
