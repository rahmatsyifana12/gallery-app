package routes

import (
	"gallery-app/controllers"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file")
	}

	jwt_access_secret := os.Getenv("JWT_ACCESS_SECRET")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	// use "g" instead of "e" if the endpoint needs an authentication middleware
	g := e.Group("", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(jwt_access_secret),
	}))

	g.POST("/memories", controllers.CreateMemory)
}