package promo

import (
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func Config(db *gorm.DB, validate *validator.Validate, router *httprouter.Router) *httprouter.Router {
	// Init Repo
	repository := NewRepositoryPromo()
	// Init Service
	service := NewServicePromo(repository, db, validate)
	// Init Controller
	controller := NewControllerPromo(service)

	router.GET("/api/v1/promos", controller.GetAllContoller)
	router.POST("/api/v1/promos", controller.CreatePromoController)

	return router
}
