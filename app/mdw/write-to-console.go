package mdw

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func WriteToConsole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Hit the page")
		return next(c)
	}
}