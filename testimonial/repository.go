package testimonial

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"gorm.io/gorm"
)

type RepositoryTestimonial interface {
	GetAll(ctx context.Context, tx *gorm.DB) []Testimonial
	Create(ctx context.Context, tx *gorm.DB, testimonial Testimonial) Testimonial
}

type repositoryTestimonial struct {
}

func NewRepositoryTestimonial() *repositoryTestimonial {
	return &repositoryTestimonial{}
}

// SQL Query Get All Testimonial
func (r *repositoryTestimonial) GetAll(ctx context.Context, tx *gorm.DB) []Testimonial {
	testimonials := []Testimonial{}
	err := tx.WithContext(ctx).Preload("User").Find(&testimonials).Error
	helper.HandleError(err)

	return testimonials
}

// SQL Query Create Testimonial
func (r *repositoryTestimonial) Create(ctx context.Context, tx *gorm.DB, testimonial Testimonial) Testimonial {
	err := tx.WithContext(ctx).Create(&testimonial).Error
	helper.HandleError(err)

	return testimonial
}
