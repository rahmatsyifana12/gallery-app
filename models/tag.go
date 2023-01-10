package models

import "time"

type Tag struct {
	ID 			uint64 `gorm:"primaryKey;autoIncrement"`
	Name 		string `gorm:"size:255;not null"`
	CreatedAt 	time.Time `gorm:"not null"`
	UpdatedAt 	time.Time `gorm:"not null"`
}

type CreateTagDTO struct {
	Name 		string `validate:"required,min=1,max=24,alphanum"`
	MemoryID 	uint64 `validate:"required"`
}