package galery

import (
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func Config(db *sql.DB, validate *validator.Validate, router *httprouter.Router) *httprouter.Router {
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
