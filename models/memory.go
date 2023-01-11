package models

import (
	"time"

	"gorm.io/gorm"
)

type Memory struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	Description string    `gorm:"size:1023;not null"`
	UserID      uint64    `gorm:"not null"`
	Tags        []*Tag    `gorm:"many2many:memory_tags;"`
	Images      []Image   `gorm:"foreignKey:MemoryID"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}

type CreateMemoryDTO struct {
	Description string `validate:"required,min=4"`
}

func GetAllMemories(db *gorm.DB, sort []string) ([]Memory, error) {
	var memories []Memory
	if sort[0] == "uploaded" {
		if sort[1] == "desc" {
			if res := db.Model(&Memory{}).Preload("Images").Preload("Tags").Order("created_at DESC").Find(&memories); res.Error != nil {
				return nil, res.Error
			}
		} else if sort[1] == "asc" {
			if res := db.Model(&Memory{}).Preload("Images").Preload("Tags").Order("created_at ASC").Find(&memories); res.Error != nil {
				return nil, res.Error
			}
		}

		return memories, nil
	} else if sort[0] == "edited" {
		if sort[1] == "desc" {
			if res := db.Model(&Memory{}).Preload("Images").Preload("Tags").Order("updated_at DESC").Find(&memories); res.Error != nil {
				return nil, res.Error
			}
		} else if sort[1] == "asc" {
			if res := db.Model(&Memory{}).Preload("Images").Preload("Tags").Order("updated_at ASC").Find(&memories); res.Error != nil {
				return nil, res.Error
			}
		}

		return memories, nil
	}

	if res := db.Model(&Memory{}).Preload("Images").Preload("Tags").Find(&memories); res.Error != nil {
		return nil, res.Error
	}

	return memories, nil
}

// func SortAllMemories(db *gorm.DB, sort []string) ([]Memory, error) {
// 	var memories []Memory

// 	if sort[0] == "uploaded" {
// 		if sort[1] == "desc" {
// 			if res := db.Model(&Memory{}).Preload("Images").Preload("Tags").Order("createdAt DESC").Find(&memories); res.Error != nil {
// 				return nil, res.Error
// 			}

// 			return memories, nil
// 		}

// 		if res := db.Model(&Memory{}).Preload("Images").Preload("Tags").Order("createdAt ASC").Find(&memories); res.Error != nil {
// 			return nil, res.Error
// 		}

// 		return memories, nil
// 	} else if sort[0] == "edited" {

// 	}
// 	return nil, errors.New("Invalid parameter")
// }

func GetMemoryByID(db *gorm.DB, ID uint64) (Memory, error) {
	var memory Memory
	if res := db.Model(&Memory{}).Preload("Images").Preload("Tags").First(&memory, ID); res.Error != nil {
		return memory, res.Error
	}

	return memory, nil
}
