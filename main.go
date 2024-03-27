package main

import (
	"band_api/controller"
	"band_api/db"
	"band_api/repository"
	"band_api/router"
	"band_api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title 	Band Service API
// @version	1.0
// @description A Band service API in Go
func main() {

	log.Info().Msg("Started server!")
	db := db.InitializeDatabase()
	validate := validator.New()

	bandRepository := repository.NewBandRepositoryImpl(db)
	bandService := service.NewBandServiceImpl(bandRepository, validate)
	bandController := controller.NewBandController(bandService)

	routes := router.NewRouter(bandController)

	server := &http.Server{
		Addr: ":8080",
	}

	http.ListenAndServe(server.Addr, routes)
}
