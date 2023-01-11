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
	e.PUT("/memories/:id", controllers.UpdateMemoryById, middlewares.Authenticate())
	e.POST("/memories/images", controllers.AddImagesToMemory, middlewares.Authenticate())
	e.PUT("/memories/images/:id", controllers.UpdateImageInMemoryById, middlewares.Authenticate())
	e.GET("/memories", controllers.GetAllMemories)
	e.GET("/memories/detail/:id", controllers.GetMemoryByID)
	e.POST("/memories/tags", controllers.AddTagsToMemory, middlewares.Authenticate())
}
