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
	"embed"
	"io/fs"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

//go:embed resources/team_photos
var resourcesTeamPhotos embed.FS

//go:embed resources/service_image
var resourcesServiceImage embed.FS

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

	// URL Access File
	directory, _ := fs.Sub(resourcesTeamPhotos, "resources/team_photos")
	router.ServeFiles("/team_photos/*filepath", http.FS(directory))
	directory1, _ := fs.Sub(resourcesServiceImage, "resources/service_image")
	router.ServeFiles("/service_image/*filepath", http.FS(directory1))

	// Handler for Exception Error
	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.HandleError(err)
}
