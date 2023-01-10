package models

import (
	"time"

	"gorm.io/gorm"
)

type Memory struct {
	ID 			uint64 `gorm:"primaryKey;autoIncrement"`
	Description string `gorm:"size:1023;not null"`
	UserID 		uint64 `gorm:"not null"`
	Tags 		[]Tag `gorm:"many2many:memory_tags;"`
	Images 		[]Image `gorm:"foreignKey:MemoryID"`
	CreatedAt 	time.Time `gorm:"not null"`
	UpdatedAt 	time.Time `gorm:"not null"`
}

type CreateMemoryDTO struct {
	Description string `validate:"required,min=4"`
}

type Tags struct {
	ID 			uint64 `gorm:"primaryKey;autoIncrement"`
	Name 		string `gorm:"size:255;not null"`
	Memory 		[]Memory `gorm:"many2many:memory_tags;"`
	CreatedAt 	time.Time `gorm:"not null"`
	UpdatedAt 	time.Time `gorm:"not null"`
}

type CreateTagsDTO struct {
	Name 		string `validate:"required,min=1,alphanum"`
}

type Images struct {
	ID 			uint64 `gorm:"primaryKey;autoIncrement"`
	Image 		string `gorm:"not null;size:1023"`
	MemoryID 	uint64 `gorm:"not null"`
	CreatedAt 	time.Time `gorm:"not null"`
	UpdatedAt 	time.Time `gorm:"not null"`
}

func GetAllMemories(db *gorm.DB) ([]Memory, error) {
	var memories []Memory
	if res := db.Model(&Memory{}).Preload("Images").Preload("Tags").Find(&memories); res.Error != nil {
		return nil, res.Error
	}

	return memories, nil
}