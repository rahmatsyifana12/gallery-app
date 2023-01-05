package main

import (
	"gallery-app/configs"
	"gallery-app/models"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file")
	}

	db := configs.DBConfig()

	// migrates all schema
	db.AutoMigrate(&models.User{})

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}