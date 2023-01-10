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
<<<<<<< HEAD
	db.AutoMigrate(&models.User{}, &models.Memory{}, &models.Tag{}, &models.Image{})
=======
	db.Migrator().DropTable(&models.Memory{}, &models.Tags{}, &models.Images{})
	db.AutoMigrate(&models.User{}, &models.Memory{}, &models.Tags{}, &models.Images{})
>>>>>>> e997337314f51bae3fed95434688fecbbe2c9839

	e := echo.New()
	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}