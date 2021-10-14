package galery

import (
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func Config(db *gorm.DB, validate *validator.Validate, router *httprouter.Router) *httprouter.Router {
	// Init Repo
	repository := NewRepositoryGalery()
	// Init Service
	service := NewServiceGalery(repository, db, validate)
	// Init Controller
	controller := NewControllerGalery(service)

	// Create Router End Point
	router.GET("/api/v1/galeries", controller.GetAllController)

	return router
}
