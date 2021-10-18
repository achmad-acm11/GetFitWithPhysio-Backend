package testimonial

import (
	"GetfitWithPhysio-backend/patient"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func Config(db *gorm.DB, validate *validator.Validate, router *httprouter.Router) *httprouter.Router {
	// Init Repo
	repository := NewRepositoryTestimonial()
	repositoryPatient := patient.NewRepositoryPatient()

	// Init Service
	service := NewServiceTestimonial(repository, repositoryPatient, db, validate)
	// Init Controller
	contoller := NewControllrerTestimonial(service)

	// Create Router End Point
	router.GET("/api/v1/testimonials", contoller.GetAllController)
	router.POST("/api/v1/testimonials", contoller.CreateTestimonialController)

	return router
}
