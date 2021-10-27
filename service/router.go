package service

import (
	"GetfitWithPhysio-backend/promo"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func Config(db *gorm.DB, validate *validator.Validate, router *httprouter.Router) *httprouter.Router {
	// Init Repository
	repository := NewRepositoryService()
	repositoryPromo := promo.NewRepositoryPromo()

	// Init Service
	service := NewServiceImpl(repository, repositoryPromo, db, validate)
	// Init Controller
	controller := NewServiceController(service)

	// Create End Point
	router.GET("/api/v1/services", controller.GetAllController)
	router.POST("/api/v1/services", controller.CreateController)
	router.GET("/api/v1/services/promo", controller.GetAllPromoController)

	return router
}
