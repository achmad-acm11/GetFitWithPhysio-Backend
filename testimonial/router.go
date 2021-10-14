package testimonial

import (
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func Config(db *gorm.DB, validate *validator.Validate, router *httprouter.Router) *httprouter.Router {
	// Init Repo
	repository := NewRepositoryTestimonial()
	// Init Service
	service := NewServiceTestimonial(repository, db, validate)
	// Init Controller
	contoller := NewControllrerTestimnoial(service)

	// Create Router End Point
	router.GET("/api/v1/testimonials", contoller.GetAllController)

	return router
}
