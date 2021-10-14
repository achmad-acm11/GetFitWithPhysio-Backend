package testimonial

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"gorm.io/gorm"
)

type RepositoryTestimonial interface {
	GetAll(ctx context.Context, tx *gorm.DB) []Testimnoial
}

type repositoryTestimonial struct {
}

func NewRepositoryTestimonial() *repositoryTestimonial {
	return &repositoryTestimonial{}
}

// SQL Query Get All Testimonial
func (r *repositoryTestimonial) GetAll(ctx context.Context, tx *gorm.DB) []Testimnoial {
	testimonials := []Testimnoial{}
	err := tx.WithContext(ctx).Find(&testimonials).Error
	helper.HandleError(err)

	return testimonials
}
