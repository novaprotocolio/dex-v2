package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler
func sayHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func HelloRoute(group *echo.Group) {
	group.GET("/", sayHello)
}
