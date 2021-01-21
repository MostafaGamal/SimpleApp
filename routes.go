package main

import "github.com/labstack/echo/v4"

// AppInterface represents all app handlers.
type AppInterface interface {
	// (POST /login)
	Login(c echo.Context) error
	// (GET /hello/{name})
	Hello(c echo.Context) error
	// (POST /request)
	Request(c echo.Context) error
}

// RegisterRoutes register all app routes
func RegisterRoutes(group *echo.Group, ai AppInterface) {
	group.POST("/login", ai.Login)
	group.GET("/hello/:name", ai.Hello)
	group.POST("/request", ai.Request)
	return
}
