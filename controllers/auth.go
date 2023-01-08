package controllers

import (
	"encoding/json"
	"gallery-app/models"
	"io/ioutil"
	"net/http"
	"time"

	"gallery-app/configs"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	user := new(models.CreateUserDTO)
	if err := c.Bind(user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// validate user input
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db := configs.DBConfig()

	// check if username already exists
	var user_exists models.User
	if err := db.First(&user_exists, "username = ?", user.Username).Error; err == nil {
		return c.String(http.StatusBadRequest, "Username already exists")
	}

	// hash raw password into hashed password
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

func Login(c echo.Context) error {
	var user models.User
	payload, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(payload, &user)

	if err != nil {
		return err
	}

	// Check if username exists
	var results models.User
	db := configs.DBConfig()
	if err := db.First(&results, "username = ?", user.Username).Error; err != nil {
		return c.String(http.StatusBadRequest, "Your credentials doens't match our records")
	}

	// Check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(results.Password), []byte(user.Password)); err != nil {
		return c.String(http.StatusBadRequest, "Your credentials doens't match our records")
	}

	//Generate JWT
	token, err := GenerateJWT(results.ID, results.Username, "secret")
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	
	var tokenJSON = map[string]string{"token": token}

	resp := c.JSON(http.StatusOK, tokenJSON)
	return resp
}

func GenerateJWT(id uint64, username string, key string) (string, error) {
	//Generate JWT for Auth
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"username": username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "tokenError", err
	}

	return tokenString, nil
}