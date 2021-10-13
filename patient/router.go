package patient

import (
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func Config(db *sql.DB, validate *validator.Validate, router *httprouter.Router) *httprouter.Router {
	// Init Repo
	repository := NewRepositoryPatient()
	// Init Service
	service := NewServicePatient(repository, db, validate)
	// Init Controller
	controller := NewControllerPatinet(service)

	// Create End Point
	router.GET("/api/v1/patients", controller.GetAllController)

	return router
}
