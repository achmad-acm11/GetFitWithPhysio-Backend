package patient

import (
	"GetfitWithPhysio-backend/user"
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func Config(db *sql.DB, validate *validator.Validate, router *httprouter.Router) *httprouter.Router {
	// Init Repo
	repository := NewRepositoryPatient()
	repositoryUser := user.NewRepositoryUser()

	// Init Service
	service := NewServicePatient(repository, repositoryUser, db, validate)
	// Init Controller
	controller := NewControllerPatinet(service)

	// Create End Point
	router.GET("/api/v1/patients", controller.GetAllController)
	router.POST("/api/v1/register", controller.Register)

	return router
}
