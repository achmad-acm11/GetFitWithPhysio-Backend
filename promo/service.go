package promo

import (
	"GetfitWithPhysio-backend/helper"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type ServicePromo interface {
	GetAllService(ctx context.Context) []PromoResponse
}

type servicePromo struct {
	repo     RepositoryPromo
	db       *sql.DB
	validate *validator.Validate
}

func NewServicePromo(repo RepositoryPromo, db *sql.DB, validate *validator.Validate) *servicePromo {
	return &servicePromo{
		repo:     repo,
		db:       db,
		validate: validate,
	}
}

func (s *servicePromo) GetAllService(ctx context.Context) []PromoResponse {
	// Start Transaction
	tx, err := s.db.Begin()
	helper.HandleError(err)
	defer helper.CommitOrRollback(tx)

	promos := s.repo.GetAll(ctx, tx)

	return MapPromosResponse(promos)
}
