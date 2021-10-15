package team

import (
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func Config(db *gorm.DB, validator *validator.Validate, router *httprouter.Router) *httprouter.Router {
	// Init Repo
	repository := NewRepository()
	// Init Service
	service := NewTeamService(repository, db, validator)
	// Init Controller
	controller := NewTeamController(service)

	// Create End Point
	router.GET("/api/v1/teams", controller.GetAllController)
	router.POST("/api/v1/teams", controller.CreateTeam)

	return router
}
