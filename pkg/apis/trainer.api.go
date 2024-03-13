package apis

import (
	"anhthi-projects/pt-booking-go/pkg/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

/*
 * Definition
 */

type TrainerApi struct {
	service services.ITrainerService
}

/*
 * Implement
 */

func (api *TrainerApi) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := api.service.Get(c.Request().Context(), id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func CreateTrainerApi(s services.ITrainerService) *TrainerApi {
	return &TrainerApi{
		service: s,
	}
}