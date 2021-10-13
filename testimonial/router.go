package testimonial

import (
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func Config(db *sql.DB, validate *validator.Validate, router *httprouter.Router) *httprouter.Router {
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
