package middlewares

import (
	"errors"
	"gallery-app/configs"
	"gallery-app/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func GetToken(c echo.Context) (*jwt.Token, error) {
	token := c.Get("user").(*jwt.Token)
	if token == nil {
		return nil, errors.New("unauthorized error")
	}
	return token, nil
}

func GetClaims(c echo.Context) (*models.JWTCustomClaims, error) {
	token, err := GetToken(c)
	if err != nil {
		return nil, errors.New("unauthorized error")
	}

	db := configs.DBConfig()
	var user models.User
	if err := db.First(&user, "token = ?", token.Raw).Error; err != nil {
		return nil, errors.New("unauthorized error")
	}

	claims := token.Claims.(*models.JWTCustomClaims)
	if claims == nil {
		return nil, errors.New("unauthorized error")
	}

	return claims, nil
}