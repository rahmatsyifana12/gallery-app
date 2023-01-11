package controllers

import (
	"gallery-app/configs"
	"gallery-app/middlewares"
	"gallery-app/models"
	"net/http"
	"strconv"

	"time"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
)

func CreateMemory(c echo.Context) error {
	memory := new(models.CreateMemoryDTO)
	if err := c.Bind(memory); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// validate user input
	validate := validator.New()
	err := validate.Struct(memory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	claims, err := middlewares.GetClaims(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	new_memory := models.Memory{
		Description: memory.Description,
		UserID:      claims.ID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	db := configs.DBConfig()
	db.Select("Description", "UserID", "CreatedAt", "UpdatedAt").Create(&new_memory)

	return c.JSON(http.StatusOK, echo.Map{
		"status":  "success",
		"message": "Succesfully created a memory",
	})
}

func UpdateMemoryById (c echo.Context) error {
	memory := new(models.CreateMemoryDTO)
	if err := c.Bind(memory); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	// validate user input
	validate := validator.New()
	err := validate.Struct(memory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	MemoryID, _ := strconv.Atoi(c.Param("id"))
	db := configs.DBConfig()
	// check if memory exists
	var memory_exists models.Memory
	if err := db.First(&memory_exists, "id = ?", MemoryID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": "Memory does not exist",
		})
	} else {
		// update memory
		db.Model(&memory_exists).Updates(models.Memory{
			Description: memory.Description,
			UpdatedAt: time.Now(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"message": "Memory successfully updated",
	})
}

func GetAllMemories(c echo.Context) error {
	db := configs.DBConfig()
	sort := strings.Split(c.Param("sort"), ".")

	memories, err := models.GetAllMemories(db, sort)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":   "success",
		"memories": memories,
	})
}

func GetMemoryByID(c echo.Context) error {
	db := configs.DBConfig()
	// Convert ID parameter to uint64
	id, errConv := strconv.ParseUint(c.Param("id"), 10, 64)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "fail",
			"message": errConv.Error(),
		})
	}

	memory, errMem := models.GetMemoryByID(db, id)

	if errMem != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "fail",
			"message": errMem.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"memory": memory,
	})
}


