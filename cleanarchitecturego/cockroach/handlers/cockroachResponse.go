package handlers

import "github.com/labstack/echo/v4"

type baseResposne struct {
	Message string `json:"message"`
}

func response(c echo.Context, responseCode int, message string) error {
	return c.JSON(responseCode, &baseResposne{
		Message: message,
	})
}
