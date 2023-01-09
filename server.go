package main

import (
	"gallery-app/configs"
	"gallery-app/models"
	"gallery-app/routes"
	"gallery-app/scripts"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file")
	}

	jwt_access_secret := os.Getenv("JWT_ACCESS_SECRET")
	if jwt_access_secret == "" {
		if err := scripts.GenerateSecret(); err != nil {
			panic(err.Error())
		}
	}

	db := configs.DBConfig()

	// migrates all schema
	db.AutoMigrate(&models.User{}, &models.Memory{}, &models.Tag{}, &models.Image{})

	e := echo.New()
	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}