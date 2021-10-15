package service

import (
	"GetfitWithPhysio-backend/exception"
	"GetfitWithPhysio-backend/helper"
	"GetfitWithPhysio-backend/promo"
	"context"
	"errors"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ServiceImpl interface {
	GetAllService(ctx context.Context) []ServiceResponse
	CreateService(ctx context.Context, req CreateServiceRequest) ServiceResponse
	UploadImageService(ctx context.Context, serviceId int, filePath string)
}

type serviceImpl struct {
	repo      RepositoryService
	repoPromo promo.RepositoryPromo
	db        *gorm.DB
	validate  *validator.Validate
}

func NewServiceImpl(repo RepositoryService, repoPromo promo.RepositoryPromo, db *gorm.DB, validate *validator.Validate) *serviceImpl {
	return &serviceImpl{
		repo:      repo,
		repoPromo: repoPromo,
		db:        db,
		validate:  validate,
	}
}

// Get All Service
func (s *serviceImpl) GetAllService(ctx context.Context) []ServiceResponse {
	// Start Transaction
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	services := s.repo.GetAll(ctx, tx)

	return MapServicesResponse(services)
}

// Create Service
func (s *serviceImpl) CreateService(ctx context.Context, req CreateServiceRequest) ServiceResponse {
	// Validate
	err := s.validate.Struct(req)
	helper.HandleError(err)

	// Start Transaction
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	if req.Kode_promo != 0 {
		promo := s.repoPromo.GetOneById(ctx, tx, req.Kode_promo)

		if promo.Id == 0 {
			panic(exception.NewNotFoundError(errors.New("kode promo not found").Error()))
		}
	}

	service := s.repo.Create(ctx, tx, Service{
		Kode_promo:   req.Kode_promo,
		Service_name: req.Service_name,
		Kuota_meet:   req.Kuota_meet,
		Price:        req.Price,
		Description:  req.Description,
	})

	return MapServiceResponse(service)
}

// Upload Image Service
func (s *serviceImpl) UploadImageService(ctx context.Context, serviceId int, filePath string) {
	// Start Transacrion
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	s.repo.UpdateImage(ctx, tx, serviceId, filePath)
}
