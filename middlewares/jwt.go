package middlewares

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func GetToken(c echo.Context) (*jwt.Token, error) {
	token := c.Get("user").(*jwt.Token)
	return token, nil
}

func GetClaims(token *jwt.Token) (jwt.Claims, error) {
	claims := token.Claims.(jwt.MapClaims)
	if claims == nil {
		return nil, errors.New("unauthorized error")
	}

	return claims, nil
}