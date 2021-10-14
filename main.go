package main

import (
	"GetfitWithPhysio-backend/app"
	"GetfitWithPhysio-backend/exception"
	"GetfitWithPhysio-backend/galery"
	"GetfitWithPhysio-backend/helper"
	"GetfitWithPhysio-backend/patient"
	"GetfitWithPhysio-backend/promo"
	"GetfitWithPhysio-backend/service"
	"GetfitWithPhysio-backend/team"
	"GetfitWithPhysio-backend/testimonial"
	"GetfitWithPhysio-backend/user"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Config Database
	db := app.ConfigDB()

	// Init Validate
	validate := validator.New()

	// Init Router
	router := httprouter.New()

	// Collection API Users
	router = user.Config(db, validate, router)
	// Collection API Teams
	router = team.Config(db, validate, router)
	// Collection API Services
	router = service.Config(db, validate, router)
	// Collection API Promos
	router = promo.Config(db, validate, router)
	// Collection API Testimonials
	router = testimonial.Config(db, validate, router)
	// Collection API Galeries
	router = galery.Config(db, validate, router)
	// Collection API Patinets
	router = patient.Config(db, validate, router)

	// Handler for Exception Error
	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.HandleError(err)
}
