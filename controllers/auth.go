package controllers

import (
	"gallery-app/entities"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gallery-app/configs"
	"github.com/go-playground/validator/v10"
)

func Register(c echo.Context) error {
	u := new(entities.CreateUserDTO)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	hashed_password, err := hashPassword(u.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	new_u := entities.User{
		Username: u.Username,
		Password: hashed_password,
		Name: u.Name,
		Phone: u.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db := configs.DBConfig()
	db.Select("Username", "Password", "Name", "Phone", "CreatedAt", "UpdatedAt").Create(&new_u)

	return c.JSON(http.StatusOK, new_u)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}