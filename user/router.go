package user

import (
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func Config(db *gorm.DB, validate *validator.Validate, router *httprouter.Router) *httprouter.Router {
	// Init Repo
	repository := NewRepositoryUser()
	// Init Service
	service := NewServiceUser(repository, db, validate)
	// Init Controller
	controller := NewControllerUser(service)

	router.POST("/api/v1/login", controller.Login)

	return router
}
