package apis

import (
	"anhthi-projects/pt-booking-go/pkg/services"
	"anhthi-projects/pt-booking-go/schemas/dtos"
	"fmt"
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

func (api *TrainerApi) List(eCtx echo.Context) error {
	trainers, err := api.service.List(eCtx.Request().Context());

	if err !=nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return eCtx.JSON(http.StatusOK, trainers)
}

func (api *TrainerApi) Get(eCtx echo.Context) error {
	id, err := strconv.Atoi(eCtx.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	trainer, err := api.service.Get(eCtx.Request().Context(), id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Trainer not found")
	}

	return eCtx.JSON(http.StatusOK, trainer)
}

func (api *TrainerApi) Create(eCtx echo.Context) error {
	trainer := dtos.TrainerDTO{}
	err := eCtx.Bind(&trainer)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	result, err := api.service.Create(eCtx.Request().Context(), trainer)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return eCtx.JSON(http.StatusOK, result)
}

func (api *TrainerApi) Update(eCtx echo.Context) error {
	id, idErr := strconv.Atoi(eCtx.Param("id"))

	if idErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	trainer := dtos.TrainerDTO{}
	bindErr := eCtx.Bind(&trainer)

	if bindErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	result, err := api.service.Update(eCtx.Request().Context(), trainer, id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return eCtx.JSON(http.StatusOK, result)
}

func (api *TrainerApi) Delete(eCtx echo.Context) error {
	id, idErr := strconv.Atoi(eCtx.Param("id"))

	if idErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

 	err := api.service.Delete(eCtx.Request().Context(), id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return eCtx.JSON(http.StatusOK, nil)
}

func CreateTrainerApi(s services.ITrainerService) *TrainerApi {
	return &TrainerApi{
		service: s,
	}
}