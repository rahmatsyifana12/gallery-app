package controllers

import (
	"gallery-app/configs"
	"gallery-app/models"
	"net/http"
	"time"
	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
)

func CreateMemory(c echo.Context) error {
	memory := new(models.CreateMemoryDTO)
	if err := c.Bind(memory); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// validate user input
	validate := validator.New()
	err := validate.Struct(memory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	new_memory := models.Memory{
		Description: memory.Description,
		UserID: 1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	db := configs.DBConfig()
	db.Select("Description", "UserID", "CreatedAt", "UpdatedAt").Create(&new_memory)

	return c.JSON(http.StatusOK, new_memory)
}

func GetAllMemories(c echo.Context) error {
	db := configs.DBConfig()

	memories, err := models.GetAllMemories(db)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "sucess",
		"data": memories,
	})
}