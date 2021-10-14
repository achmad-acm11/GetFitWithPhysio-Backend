package service

import (
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func Config(db *gorm.DB, validate *validator.Validate, router *httprouter.Router) *httprouter.Router {
	// Init Repository
	repository := NewRepositoryService()
	// Init Service
	service := NewServiceImpl(repository, db, validate)
	// Init Controller
	controller := NewServiceController(service)

	// Create End Point
	router.GET("/api/v1/services", controller.GetAllController)

	return router
}
