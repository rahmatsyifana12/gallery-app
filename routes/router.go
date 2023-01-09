package routes

import (
	"gallery-app/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	e.POST("/memories", controllers.CreateMemory)
}