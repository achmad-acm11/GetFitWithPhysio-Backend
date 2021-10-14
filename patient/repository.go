package patient

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"gorm.io/gorm"
)

type ReposioryPatient interface {
	GetAll(ctx context.Context, tx *gorm.DB) []Patient
	Create(ctx context.Context, tx *gorm.DB, patient Patient) Patient
}

type repositoryPatient struct {
}

func NewRepositoryPatient() *repositoryPatient {
	return &repositoryPatient{}
}

func (r *repositoryPatient) GetAll(ctx context.Context, tx *gorm.DB) []Patient {
	patients := []Patient{}
	err := tx.WithContext(ctx).Preload("User").Find(&patients).Error
	helper.HandleError(err)

	return patients
}

func (r *repositoryPatient) Create(ctx context.Context, tx *gorm.DB, patient Patient) Patient {
	err := tx.WithContext(ctx).Create(&patient).Error
	helper.HandleError(err)

	return patient
}
