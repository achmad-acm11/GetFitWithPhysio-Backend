package main

import (
	"GetfitWithPhysio-backend/app"
	"GetfitWithPhysio-backend/helper"
	"GetfitWithPhysio-backend/service"
	"GetfitWithPhysio-backend/team"
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

	// Collection API Teams
	router = team.Config(db, validate, router)
	// Collection API Services
	router = service.Config(db, validate, router)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.HandleError(err)
}
