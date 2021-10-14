package patient

import (
	"GetfitWithPhysio-backend/user"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func Config(db *gorm.DB, validate *validator.Validate, router *httprouter.Router) *httprouter.Router {
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
	router.POST("/api/v1/patients", controller.Create)

	return router
}
