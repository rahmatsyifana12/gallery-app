package controllers

import (
	"gallery-app/configs"
	"gallery-app/models"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"time"

	"github.com/labstack/echo/v4"
)

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