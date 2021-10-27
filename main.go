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
	"embed"
	"io/fs"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

//go:embed resources/team_photos
var resourcesTeamPhotos embed.FS

//go:embed resources/service_image
var resourcesServiceImage embed.FS

//go:embed resources/galeries
var resourcesGaleries embed.FS

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
	directory, _ := fs.Sub(resourcesTeamPhotos, "resources/team_photos")
	router.ServeFiles("/team_photos/*filepath", http.FS(directory))
	directory1, _ := fs.Sub(resourcesServiceImage, "resources/service_image")
	router.ServeFiles("/service_image/*filepath", http.FS(directory1))
	directory2, _ := fs.Sub(resourcesGaleries, "resources/galeries")
	router.ServeFiles("/galery/*filepath", http.FS(directory2))

	// Handler for Exception Error
	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: ":" + os.Getenv("PORT"),
		// Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.HandleError(err)
}
