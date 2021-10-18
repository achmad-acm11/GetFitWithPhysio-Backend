package galery

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ServiceGalery interface {
	GetAllService(ctx context.Context) []GaleryResponse
	CreateService(ctx context.Context, req CreateGaleryRequest) GaleryResponse
	UploadPhoto(ctx context.Context, galeryId int, filePath string)
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

// Get All Galeries Service
func (s *serviceGalery) GetAllService(ctx context.Context) []GaleryResponse {
	// Start Transaction
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	galeries := s.repo.GetAll(ctx, tx)

	return MapGaleriesResponse(galeries)
}

func (s *serviceGalery) CreateService(ctx context.Context, req CreateGaleryRequest) GaleryResponse {
	err := s.validate.Struct(req)
	helper.HandleError(err)

	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	galery := s.repo.Create(ctx, tx, Galery{
		Caption:    req.Caption,
		SubCaption: req.SubCaption,
	})

	return MapGaleryResponse(galery)
}
func (s *serviceGalery) UploadPhoto(ctx context.Context, galeryId int, filePath string) {
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	s.repo.UploadPhoto(ctx, tx, galeryId, filePath)
}
