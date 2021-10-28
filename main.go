package main

import (
	"GetfitWithPhysio-backend/app"
	"GetfitWithPhysio-backend/appointment"
	"GetfitWithPhysio-backend/exception"
	"GetfitWithPhysio-backend/galery"
	"GetfitWithPhysio-backend/helper"
	"GetfitWithPhysio-backend/patient"
	"GetfitWithPhysio-backend/promo"
	"GetfitWithPhysio-backend/service"
	"GetfitWithPhysio-backend/team"
	"GetfitWithPhysio-backend/testimonial"
	"GetfitWithPhysio-backend/transaction"
	"GetfitWithPhysio-backend/user"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
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
	// Collection API Transaction
	router = transaction.Config(db, validate, router)
	// Collection API Appointment
	router = appointment.Config(db, validate, router)

	// URL Access File
	// directory, _ := fs.Sub(resourcesTeamPhotos, "resources/team_photos")
	router.ServeFiles("/team_photos/*filepath", http.Dir("resources/team_photos"))
	// directory1, _ := fs.Sub(resourcesServiceImage, "resources/service_image")
	router.ServeFiles("/service_image/*filepath", http.Dir("resources/service_image"))
	// directory2, _ := fs.Sub(resourcesGaleries, "resources/galeries")
	router.ServeFiles("/galery/*filepath", http.Dir("resources/galeries"))

	// Handler for Exception Error
	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: ":" + os.Getenv("PORT"),
		// Addr:    "localhost:3000",
		Handler: cors.Default().Handler(router),
	}

	err := server.ListenAndServe()
	helper.HandleError(err)
}
