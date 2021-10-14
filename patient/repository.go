package patient

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"gorm.io/gorm"
)

type ReposioryPatient interface {
	GetAll(ctx context.Context, tx *gorm.DB) []Patient
	Create(ctx context.Context, tx *gorm.DB, patient Patient) Patient
	GetOneById(ctx context.Context, tx *gorm.DB, patientId int) Patient
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

func (r *repositoryPatient) GetOneById(ctx context.Context, tx *gorm.DB, patientId int) Patient {
	patient := Patient{}

	err := tx.WithContext(ctx).Preload("User").Where("id = ?", patientId).Find(&patient).Error
	helper.HandleError(err)

	return patient
}
