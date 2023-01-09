package middlewares

import (
	"gallery-app/models"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Authenticate() echo.MiddlewareFunc {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file")
	}
	jwt_access_secret := os.Getenv("JWT_ACCESS_SECRET")

	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(jwt_access_secret),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JWTCustomClaims)
		},
	})
}