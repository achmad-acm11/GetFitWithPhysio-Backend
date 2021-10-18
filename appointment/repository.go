package appointment

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"gorm.io/gorm"
)

type RepositoryAppointment interface {
	GetAll(cix context.Context, tx *gorm.DB) []Appointment
	GetOneById(ctx context.Context, tx *gorm.DB, appointmentId int) Appointment
	Create(ctx context.Context, tx *gorm.DB, appointment Appointment) Appointment
}

type repositoryAppontment struct {
}

func NewRepositoryAppointment() *repositoryAppontment {
	return &repositoryAppontment{}
}

// Query SQL Get All Data Appointment
func (r *repositoryAppontment) GetAll(ctx context.Context, tx *gorm.DB) []Appointment {
	appointment := []Appointment{}

	err := tx.WithContext(ctx).Preload("Patient").Preload("Service").Find(&appointment).Error
	helper.HandleError(err)

	return appointment
}

// Query SQL Get One Data Appointment
func (r *repositoryAppontment) GetOneById(ctx context.Context, tx *gorm.DB, appointmentId int) Appointment {
	appointment := Appointment{}

	err := tx.WithContext(ctx).Preload("Patient").Preload("Service").Where("id = ?", appointmentId).Find(&appointment).Error
	helper.HandleError(err)

	return appointment
}

// Query SQL Create Appointment
func (r *repositoryAppontment) Create(ctx context.Context, tx *gorm.DB, appointment Appointment) Appointment {
	err := tx.WithContext(ctx).Create(&appointment).Error

	helper.HandleError(err)

	return appointment
}
