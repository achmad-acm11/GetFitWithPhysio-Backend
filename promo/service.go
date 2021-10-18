package promo

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ServicePromo interface {
	GetAllService(ctx context.Context) []PromoResponse
	CreatePromoService(ctx context.Context, req CreatePromoRequest) PromoResponse
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

// Get All Promo
func (s *servicePromo) GetAllService(ctx context.Context) []PromoResponse {
	// Start Transaction
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	promos := s.repo.GetAll(ctx, tx)

	return MapPromosResponse(promos)
}

// Create Promo Service
func (s *servicePromo) CreatePromoService(ctx context.Context, req CreatePromoRequest) PromoResponse {
	// Validate
	err := s.validate.Struct(req)
	helper.HandleError(err)

	// Start Transaction
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	promo := s.repo.Create(ctx, tx, Promo{
		Discount: req.Discount,
	})

	return MapPromoResponse(promo)
}
