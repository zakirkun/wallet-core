package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter() http.Handler {

	r := echo.New()

	r.GET("/", func(c echo.Context) error {
		type ResponseOke struct {
			Message string
		}

		result := ResponseOke{Message: "Success"}

		return c.JSON(http.StatusOK, result)
	})

	return r
}
