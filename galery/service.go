package galery

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ServiceGalery interface {
	GetAllService(ctx context.Context) []GaleryResponse
}
type serviceGalery struct {
	repo     RepositoryGalery
	db       *gorm.DB
	validate *validator.Validate
}

func NewServiceGalery(repo RepositoryGalery, db *gorm.DB, validate *validator.Validate) *serviceGalery {
	return &serviceGalery{
		repo:     repo,
		db:       db,
		validate: validate,
	}
}

func (s *serviceGalery) GetAllService(ctx context.Context) []GaleryResponse {
	// Start Transaction
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	galeries := s.repo.GetAll(ctx, tx)

	return MapGaleriesResponse(galeries)
}
