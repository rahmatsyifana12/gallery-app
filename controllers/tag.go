package controllers

import (
	"gallery-app/configs"
	"gallery-app/models"
	"net/http"

	"time"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
)

func AddTagsToMemory(c echo.Context) error {
	tags := new(models.CreateTagDTO)
	if err := c.Bind(tags); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// validate user input
	validate := validator.New()
	err := validate.Struct(tags)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	db := configs.DBConfig()
	// check if tags already exist
	var tags_exists models.Tag
	if err := db.First(&tags_exists, "name = ?", tags.Name).Error; err != nil {
		// if tags don't exist exists, create tags
		new_tags := models.Tag{
			Name:      tags.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		db.Select("Name", "Created_at", "Updated_at").Create(&new_tags)
	}

	db.First(&tags_exists, "name = ?", tags.Name)
	
	new_memory_tags := models.MemoryTag {
		TagID: tags_exists.ID,
		MemoryID: tags.MemoryID,
	}
	if res := db.Select("TagID", "MemoryID").Create(&new_memory_tags); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": "fail",
			"message": "Duplicate tags in memory detected",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  "success",
		"message": "Tags successfully added",
	})
}
