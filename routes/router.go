package routes

import (
	"gallery-app/controllers"
	"gallery-app/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)
	e.PATCH("/logout", controllers.Logout, middlewares.Authenticate())

	e.POST("/memories", controllers.CreateMemory, middlewares.Authenticate())
}