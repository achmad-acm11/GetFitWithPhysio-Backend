package testimonial

import (
	"GetfitWithPhysio-backend/helper"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type ServiceTestimonial interface {
	GetAllService(ctx context.Context) []TestimonialResponse
}

type serviceTestimonial struct {
	repo      RepositoryTestimonial
	db        *sql.DB
	validator *validator.Validate
}

func NewServiceTestimonial(repo RepositoryTestimonial, db *sql.DB, validate *validator.Validate) *serviceTestimonial {
	return &serviceTestimonial{
		repo:      repo,
		db:        db,
		validator: validate,
	}
}

func (s *serviceTestimonial) GetAllService(ctx context.Context) []TestimonialResponse {
	// Start Transaction
	tx, err := s.db.Begin()
	helper.HandleError(err)

	helper.CommitOrRollback(tx)

	testimonials := s.repo.GetAll(ctx, tx)

	return MapTestimonialsResponse(testimonials)
}
