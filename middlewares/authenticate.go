package middlewares

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

func Authenticate(c echo.Context) {
	auth := c.Request().Header.Get("Authorization")
	if auth == "" {
		panic("You are unauthorized")
	}
	token := strings.Split(auth, " ")[1]
	fmt.Println(token)
}