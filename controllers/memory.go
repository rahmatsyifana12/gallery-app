package controllers

import (
	"gallery-app/configs"
	"gallery-app/middlewares"
	"gallery-app/models"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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

func AddImagesToMemory(c echo.Context) error {
	// Initialize the multipartform
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	// Get the files from the request
	files := form.File["Image"]
	isFile := false
	for _, file := range files {
		// Validate file extension
		if !(validateFileExt(strings.ToLower(file.Filename))) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"status":  "fail",
				"message": "File must be either a .jpg, .jpeg, .png",
			})
		}

		// Validate file size max 5MB
		if file.Size > int64(5000000) && file.Size < int64(0) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"status":  "fail",
				"message": "Images can't be more than 5MB",
			})
		}
	}
	for _, file := range files {
		// Create fileName with format
		var t = time.Now()
		var fileName = t.Format("20060102150405_") + RandStringBytes(16) + filepath.Ext(strings.ToLower(file.Filename))

		db := configs.DBConfig()
		MemoryID := c.FormValue("MemoryID")
		typeCastedMemoryID, err := strconv.ParseUint(string(MemoryID), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		new_image := models.Image{
			Image:     fileName,
			MemoryID:  typeCastedMemoryID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Select("Image", "MemoryID", "CreatedAt", "UpdatedAt").Create(&new_image).Error; err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		// Get file source
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}
		defer src.Close()

		// Get file destination
		dst, err := os.Create(filepath.Join("storage/images/", fileName))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}
		defer dst.Close()

		// Copy files to storage/images
		if _, err = io.Copy(dst, src); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		} else {
			isFile = true
		}
	}

	if isFile {
		return c.JSON(http.StatusOK, echo.Map{
			"status":  "success",
			"message": "Images successfully uploaded",
		})
	}
	return c.JSON(http.StatusBadRequest, echo.Map{
		"status":  "fail",
		"message": "bad request",
	})
}

func UpdateImageInMemoryById (c echo.Context) error {
	// Get file from request
	file, err := c.FormFile("Image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}
	isFile := false
	// Validate file extension
	if !(validateFileExt(strings.ToLower(file.Filename))) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": "File must be either a .jpg, .jpeg, .png",
		})
	}

	// Validate file size max 5MB
	if file.Size > int64(5000000) && file.Size < int64(0) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": "Images can't be more than 5MB",
		})
	}

	ImageID, _ := strconv.Atoi(c.Param("id"))
	db := configs.DBConfig()
	
	// Create new fileName with format
	var t = time.Now()
	var fileName = t.Format("20060102150405_")+RandStringBytes(16)+filepath.Ext(strings.ToLower(file.Filename))

	// Get file source
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}
	defer src.Close()
	
	// Get file destination
	dst, err := os.Create(filepath.Join("storage/images/", fileName))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}
	defer dst.Close()

	// Copy files to storage/images
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "fail",
			"message": err.Error(),
		})
	} else {
		isFile = true
	}

	// Update image
	if (isFile) {
		var image models.Image
		if err := db.First(&image, "id = ?", ImageID).Error; err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"status": "fail",
				"message": "Image does not exist",
			})
		} else {
			// Delete existing image
			os.Remove(filepath.Join("storage/images/", image.Image))

			// Update image
			db.Model(&image).Updates(models.Image{
				Image: fileName,
				UpdatedAt: time.Now(),
			})
		}
	}


	if (isFile) {
		return c.JSON(http.StatusOK, echo.Map{
			"status": "success",
			"message": "Images successfully updated",
		})
	}
	return c.JSON(http.StatusBadRequest, echo.Map{
		"status": "fail",
		"message": "bad request",
	})
}

func validateFileExt(s string) bool {
	fileExts := []string{".jpg", ".jpeg", ".png"}
	for _, fileExt := range fileExts {
		if strings.HasSuffix(s, fileExt) {
			return true
		}
	}
	return false
}

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GetAllMemories(c echo.Context) error {
	db := configs.DBConfig()

	memories, err := models.GetAllMemories(db)
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

