package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *app) Hello(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (app *app) Request(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}