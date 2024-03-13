package main

import (
	"anhthi-projects/pt-booking-go/app"
	"anhthi-projects/pt-booking-go/pkg/apis"
	dataaccess "anhthi-projects/pt-booking-go/pkg/data-access"
	"anhthi-projects/pt-booking-go/pkg/services"
	"net/http"

	"github.com/labstack/echo/v4"
)


func main() {
	sqlDB := app.InitDatabase("postgres://user:password@localhost:5432/pt-booking?sslmode=disable")

	// Data access
	trainerDA := dataaccess.CreateTrainerDA(sqlDB)

	// Services
	trainerService := services.CreateTrainerService(trainerDA)

	// APIs
	trainerApi := apis.CreateTrainerApi(trainerService)

	/*
	 * Routing
	 */

	server := app.InitServer()

	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "It works!")
	})

	trainer := server.Group("/trainer");
	trainer.GET("/:id", trainerApi.Get)

	server.Logger.Fatal(server.Start("127.0.0.1:3000"))
}