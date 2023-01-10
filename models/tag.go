package models

import "time"

type Tag struct {
	ID 			uint64 `gorm:"primaryKey;autoIncrement"`
	Name 		string `gorm:"size:255;not null"`
	Memory 		[]Memory `gorm:"many2many:memory_tags;"`
	CreatedAt 	time.Time `gorm:"not null"`
	UpdatedAt 	time.Time `gorm:"not null"`
}

type CreateTagDTO struct {
	Name 		string `validate:"required,min=1,alphanum"`
}