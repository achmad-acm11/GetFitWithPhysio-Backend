package service

import (
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func Config(db *sql.DB, validate *validator.Validate, router *httprouter.Router) *httprouter.Router {
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
