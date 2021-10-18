package appointment

import (
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func Config(db *gorm.DB, validate *validator.Validate, router *httprouter.Router) *httprouter.Router {
	// Init Repo
	repository := NewRepositoryAppointment()

	// Init Service
	service := NewServiceAppoinment(repository, db, validate)

	// Init Controller
	controller := NewControllerAppointment(service)

	// Crate End Point API
	router.GET("/api/v1/appointments", controller.GetAllController)
	router.POST("/api/v1/appointments/:serviceId", controller.CreateAppointment)
	router.GET("/api/v1/appointments/:appointmentId", controller.DetailAppointment)

	return router

}
