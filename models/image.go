package models

import "time"

type Image struct {
	ID 			uint64 `gorm:"primaryKey;autoIncrement"`
	Image 		string `gorm:"size:1023;not null"`
	MemoryID 	uint64 `gorm:"not null"`
	CreatedAt 	time.Time `gorm:"not null"`
	UpdatedAt 	time.Time `gorm:"not null"`
}

type CreateImageDTO struct {
	Image 		string `validate:"required,min=1,alphanum"`
}