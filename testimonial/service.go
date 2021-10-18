package testimonial

import (
	"GetfitWithPhysio-backend/exception"
	"GetfitWithPhysio-backend/helper"
	"GetfitWithPhysio-backend/patient"
	"context"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ServiceTestimonial interface {
	GetAllService(ctx context.Context) []TestimonialResponse
	CreateTestimonialService(ctx context.Context, req RequestTestimonial) TestimonialResponse
}

type serviceTestimonial struct {
	repo        RepositoryTestimonial
	repoPatient patient.ReposioryPatient
	db          *gorm.DB
	validator   *validator.Validate
}

func NewServiceTestimonial(repo RepositoryTestimonial, repoPatient patient.ReposioryPatient, db *gorm.DB, validate *validator.Validate) *serviceTestimonial {
	return &serviceTestimonial{
		repo:        repo,
		repoPatient: repoPatient,
		db:          db,
		validator:   validate,
	}
}

// Get All Data Testimonial Service
func (s *serviceTestimonial) GetAllService(ctx context.Context) []TestimonialResponse {
	// Start Transaction
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	testimonials := s.repo.GetAll(ctx, tx)

	return MapTestimonialsResponse(testimonials)
}

// Create Testimonial Service
func (s *serviceTestimonial) CreateTestimonialService(ctx context.Context, req RequestTestimonial) TestimonialResponse {
	// Check Validate
	err := s.validator.Struct(req)
	helper.HandleError(err)

	// Start Transaction
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	patient := s.repoPatient.GetOneById_user(ctx, tx, req.Id_user)

	if patient.Id == 0 {
		panic(exception.NewNotFoundError("User Not Found"))
	}

	testimonial := s.repo.Create(ctx, tx, Testimonial{
		Id_user: req.Id_user,
		Content: req.Content,
	})

	testimonial.User.Photo_user = patient.User.Photo_user
	testimonial.User.Name = patient.User.Name
	testimonial.User.Email = patient.User.Email

	return MapTestimonialResponse(testimonial)
}
