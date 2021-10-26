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
	result := []map[string]interface{}{}
	err := tx.Table("testimonials").Joins("JOIN users ON users.id = testimonials.id_user").Joins("JOIN patients ON patients.id_user = users.id").Find(&result).Error
	helper.HandleError(err)

	testimonials := MapTestimonials(result)

	return testimonials
}

// SQL Query Create Testimonial
func (r *repositoryTestimonial) Create(ctx context.Context, tx *gorm.DB, testimonial Testimonial) Testimonial {
	err := tx.WithContext(ctx).Create(&testimonial).Error
	helper.HandleError(err)

	return testimonial
}
