package promo

import (
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func Config(db *sql.DB, validate *validator.Validate, router *httprouter.Router) *httprouter.Router {
	// Init Repo
	repository := NewRepositoryPromo()
	// Init Service
	service := NewServicePromo(repository, db, validate)
	// Init Controller
	controller := NewControllerPromo(service)

	router.GET("/api/v1/promos", controller.GetAllContoller)

	return router
}
