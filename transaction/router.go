package transaction

import (
	"GetfitWithPhysio-backend/patient"
	"GetfitWithPhysio-backend/service"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func Config(db *gorm.DB, validate *validator.Validate, router *httprouter.Router) *httprouter.Router {
	// Init Repo
	respository := NewRepositoryTransaction()
	repositoryService := service.NewRepositoryService()
	repositoryPatient := patient.NewRepositoryPatient()

	// Init Service
	service := NewServiceTransaction(respository, repositoryService, repositoryPatient, db, validate)

	// Init Controller
	contoller := NewControllerTransaction(service)

	router.GET("/api/v1/transactions", contoller.GetAllTransactionController)
	router.POST("/api/v1/transactions/:serviceId", contoller.CreateTransactionController)

	return router

}
