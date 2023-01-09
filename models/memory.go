package models

import "time"

type Memory struct {
	ID 			uint64 `gorm:"primaryKey;autoIncrement"`
	Description string `gorm:"size:1023;not null"`
	UserID 		uint64 `gorm:"not null"`
	Tags 		[]Tag `gorm:"many2many:memory_tags;"`
	Images 		[]Images `gorm:"foreignKey:MemoryID"`
	CreatedAt 	time.Time `gorm:"not null"`
	UpdatedAt 	time.Time `gorm:"not null"`
}

type CreateMemoryDTO struct {
	Description string `validate:"required,min=4,alphanum"`
}
