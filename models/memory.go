package models

import "time"

type Memory struct {
	ID 			uint64 `gorm:"primaryKey;autoIncrement"`
	Description string `gorm:"not null"`
	UserID 		uint64 `gorm:"not null"`
	Tags 		[]Tags `gorm:"many2many:memory_tags;"`
	Images 		[]Images `gorm:"foreignKey:MemoryID"`
	CreatedAt 	time.Time `gorm:"not null"`
	UpdatedAt 	time.Time `gorm:"not null"`
}

type CreateMemoryDTO struct {
	Description string `validate:"required,min=4,alphanum"`
}

type Tags struct {
	ID 			uint64 `gorm:"primaryKey;autoIncrement"`
	Name 		string `gorm:"not null"`
	Memory 		[]Memory `gorm:"many2many:memory_tags;"`
	CreatedAt 	time.Time `gorm:"not null"`
	UpdatedAt 	time.Time `gorm:"not null"`
}

type CreateTagsDTO struct {
	Name 		string `validate:"required,min=1,alphanum"`
}

type Images struct {
	ID 			uint64 `gorm:"primaryKey;autoIncrement"`
	Image 		string `gorm:"not null"`
	MemoryID 	uint64 `gorm:"not null"`
	CreatedAt 	time.Time `gorm:"not null"`
	UpdatedAt 	time.Time `gorm:"not null"`
}

type CreateImagesDTO struct {
	Image 		string `validate:"required,min=1,alphanum"`
}

