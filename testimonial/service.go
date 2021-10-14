package testimonial

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ServiceTestimonial interface {
	GetAllService(ctx context.Context) []TestimonialResponse
}

type serviceTestimonial struct {
	repo      RepositoryTestimonial
	db        *gorm.DB
	validator *validator.Validate
}

func NewServiceTestimonial(repo RepositoryTestimonial, db *gorm.DB, validate *validator.Validate) *serviceTestimonial {
	return &serviceTestimonial{
		repo:      repo,
		db:        db,
		validator: validate,
	}
}

func (s *serviceTestimonial) GetAllService(ctx context.Context) []TestimonialResponse {
	// Start Transaction
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	testimonials := s.repo.GetAll(ctx, tx)

	return MapTestimonialsResponse(testimonials)
}
