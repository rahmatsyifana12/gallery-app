package middlewares

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetToken(c echo.Context) (*jwt.Token, error) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil, errors.New("Unauthorized")
	}

	return token, nil
}

func GetClaims(token *jwt.Token) (jwt.Claims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Unauthorized")
	}

	return claims, nil
}