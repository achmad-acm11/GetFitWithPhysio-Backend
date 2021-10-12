package team

import (
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func Config(db *sql.DB, validator *validator.Validate, router *httprouter.Router) *httprouter.Router {
	// Init Repo
	repository := NewRepository()
	// Init Service
	service := NewTeamService(repository, db, validator)
	// Init Controller
	controller := NewTeamController(service)

	// Create End Point
	router.GET("/api/v1/teams", controller.GetAllController)

	return router
}
