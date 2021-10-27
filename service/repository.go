package service

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"gorm.io/gorm"
)

type RepositoryService interface {
	GetAll(ctx context.Context, tx *gorm.DB) []Service
	GetAllPromo(ctx context.Context, tx *gorm.DB) []Service
	GetOneById(ctx context.Context, tx *gorm.DB, serviceId int) Service
	Create(ctx context.Context, tx *gorm.DB, service Service) Service
	UpdateImage(ctx context.Context, tx *gorm.DB, serviceId int, filePath string)
}

type repositoryService struct {
}

func NewRepositoryService() *repositoryService {
	return &repositoryService{}
}

// SQL Query Get All Data Service
func (r *repositoryService) GetAll(ctx context.Context, tx *gorm.DB) []Service {
	services := []Service{}
	err := tx.WithContext(ctx).Preload("Promo").Find(&services).Error
	helper.HandleError(err)

	return services
}

// SQL Query Get All Data Service
func (r *repositoryService) GetAllPromo(ctx context.Context, tx *gorm.DB) []Service {
	services := []Service{}
	err := tx.WithContext(ctx).Preload("Promo").Where("kode_promo != ?", 0).Find(&services).Error
	helper.HandleError(err)

	return services
}

// SQL Query Get One Service
func (r *repositoryService) GetOneById(ctx context.Context, tx *gorm.DB, serviceId int) Service {
	service := Service{}

	err := tx.WithContext(ctx).Preload("Promo").Where("id = ?", serviceId).Find(&service).Error

	helper.HandleError(err)

	return service
}

// SQL Query Create Service
func (r *repositoryService) Create(ctx context.Context, tx *gorm.DB, service Service) Service {
	err := tx.WithContext(ctx).Preload("Promo").Create(&service).Error
	helper.HandleError(err)

	return service
}

// SQL Query Upload Image Service
func (r *repositoryService) UpdateImage(ctx context.Context, tx *gorm.DB, serviceId int, filePath string) {
	err := tx.WithContext(ctx).Model(Service{}).Where("id = ?", serviceId).Update("image", filePath).Error
	helper.HandleError(err)
}
