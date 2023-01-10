package models

import (
	"time"

	"gorm.io/gorm"
)

type Memory struct {
	ID 			uint64 `gorm:"primaryKey;autoIncrement"`
	Description string `gorm:"size:1023;not null"`
	UserID 		uint64 `gorm:"not null"`
	Tags 		[]*Tag `gorm:"many2many:memory_tags;"`
	Images 		[]Image `gorm:"foreignKey:MemoryID"`
	CreatedAt 	time.Time `gorm:"not null"`
	UpdatedAt 	time.Time `gorm:"not null"`
}

type CreateMemoryDTO struct {
	Description string `validate:"required,min=4"`
}

func GetAllMemories(db *gorm.DB) ([]Memory, error) {
	var memories []Memory
	if res := db.Model(&Memory{}).Preload("Images").Preload("Tags").Find(&memories); res.Error != nil {
		return nil, res.Error
	}

	return memories, nil
}