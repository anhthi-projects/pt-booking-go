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
	sqlDB := app.InitDatabase("postgres://user:password@localhost:5432/pt-booking-service?sslmode=disable")

	/*
	* Data access
	*/

	trainerDA := dataaccess.CreateTrainerDA(sqlDB)

	/*
	* Services
	*/

	trainerService := services.CreateTrainerService(trainerDA)

	/*
	* Apis
	*/

	trainerApi := apis.CreateTrainerApi(trainerService)

	/*
	 * Routing
	 */

	server := app.InitServer()

	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "It works!")
	})

	trainer := server.Group("/trainer");
	trainer.GET("/all", trainerApi.List)
	trainer.GET("/:id", trainerApi.Get)
	trainer.POST("", trainerApi.Create)
	trainer.PUT("/:id", trainerApi.Update)
	trainer.DELETE("/:id", trainerApi.Delete)

	server.Logger.Fatal(server.Start("127.0.0.1:4000"))
}