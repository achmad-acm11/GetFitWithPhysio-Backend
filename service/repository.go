package service

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"gorm.io/gorm"
)

type RepositoryService interface {
	GetAll(ctx context.Context, tx *gorm.DB) []Service
}

type repositoryService struct {
}

func NewRepositoryService() *repositoryService {
	return &repositoryService{}
}

func (r *repositoryService) GetAll(ctx context.Context, tx *gorm.DB) []Service {
	services := []Service{}
	err := tx.WithContext(ctx).Find(&services).Error
	helper.HandleError(err)

	return services
}
