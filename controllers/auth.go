package controllers

import (
	"encoding/json"
	"gallery-app/middlewares"
	"gallery-app/models"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"gallery-app/configs"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	user := new(models.CreateUserDTO)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	// validate user input
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	db := configs.DBConfig()

	// check if username already exists
	var user_exists models.User
	if err := db.First(&user_exists, "username = ?", user.Username).Error; err == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": "Username already exists",
		})
	}

	// hash raw password into hashed password
	hashed_password, err := hashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	new_user := models.User{
		Username: user.Username,
		Password: hashed_password,
		Name: user.Name,
		Phone: user.Phone,
		Token: "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db.Select("Username", "Password", "Name", "Phone", "CreatedAt", "UpdatedAt").Create(&new_user)

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"message": "Succesfully created a new account",
	})
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
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	// Check if username exists
	var results models.User
	db := configs.DBConfig()
	if err := db.First(&results, "username = ?", user.Username).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": "Your credentials doens't match our records",
		})
	}

	err = godotenv.Load()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": "fail",
			"message": "Internal server error",
		})
	}

	// Check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(results.Password), []byte(user.Password)); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": "Your credentials doens't match our records",
		})
	}

	// Generate JWT
	jwt_access_secret := os.Getenv("JWT_ACCESS_SECRET")
	token, err := GenerateJWT(results.ID, results.Username, jwt_access_secret)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	results.Token = token
	db.Save(&results)

	return c.JSON(http.StatusOK, echo.Map{
		"status": "suscess",
		"message": "Successfully logged in",
		"token": token,
	})
}

func GenerateJWT(id uint64, username string, key string) (string, error) {
	// Set custom claims with id and username
	claims := models.JWTCustomClaims{
		ID: id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	
	// Return token
	return t, nil
}

func Logout(c echo.Context) error {
	claims, err := middlewares.GetClaims(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	db := configs.DBConfig()
	var user models.User
	if err := db.First(&user, "id = ?", claims.ID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	user.Token = ""
	db.Save(&user)
	
	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"message": "Successfully logged out",
	})
}
