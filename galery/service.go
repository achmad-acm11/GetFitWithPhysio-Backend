package galery

import (
	"GetfitWithPhysio-backend/helper"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type ServiceGalery interface {
	GetAllService(ctx context.Context) []GaleryResponse
}
type serviceGalery struct {
	repo     RepositoryGalery
	db       *sql.DB
	validate *validator.Validate
}

func NewServiceGalery(repo RepositoryGalery, db *sql.DB, validate *validator.Validate) *serviceGalery {
	return &serviceGalery{
		repo:     repo,
		db:       db,
		validate: validate,
	}
}

func (s *serviceGalery) GetAllService(ctx context.Context) []GaleryResponse {
	// Start Transaction
	tx, err := s.db.Begin()
	helper.HandleError(err)
	defer helper.CommitOrRollback(tx)

	galeries := s.repo.GetAll(ctx, tx)

	return MapGaleriesResponse(galeries)
}
