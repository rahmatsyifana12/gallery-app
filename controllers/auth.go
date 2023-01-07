package controllers

import (
	"gallery-app/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gallery-app/configs"
	"github.com/go-playground/validator/v10"
)

func Register(c echo.Context) error {
	user := new(models.CreateUserDTO)
	if err := c.Bind(user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	hashed_password, err := hashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	new_user := models.User{
		Username: user.Username,
		Password: hashed_password,
		Name: user.Name,
		Phone: user.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db := configs.DBConfig()
	db.Select("Username", "Password", "Name", "Phone", "CreatedAt", "UpdatedAt").Create(&new_user)

	return c.JSON(http.StatusOK, new_user)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}