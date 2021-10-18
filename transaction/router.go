package transaction

import (
	"GetfitWithPhysio-backend/service"
	"GetfitWithPhysio-backend/user"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func Config(db *gorm.DB, validate *validator.Validate, router *httprouter.Router) *httprouter.Router {
	// Init Repo
	respository := NewRepositoryTransaction()
	repositoryService := service.NewRepositoryService()
	repositoryUser := user.NewRepositoryUser()

	// Init Service
	service := NewServiceTransaction(respository, repositoryService, repositoryUser, db, validate)

	// Init Controller
	contoller := NewControllerTransaction(service)

	router.POST("/api/v1/transactions/:serviceId", contoller.CreateTransactionController)

	return router

}
