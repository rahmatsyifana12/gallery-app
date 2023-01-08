package models

import "time"

type User struct {
	ID 			uint64 `gorm:"primaryKey;autoIncrement"`
	Username 	string `gorm:"size:255;not null"`
	Password 	string `gorm:"size:255;not null"`
	Name 		string `gorm:"size:255;not null"`
	Phone 		string `gorm:"size:63"`
	Memory 		[]Memory `gorm:"foreignKey:UserID"`
	CreatedAt 	time.Time `gorm:"not null"`
	UpdatedAt 	time.Time `gorm:"not null"`
}

type CreateUserDTO struct {
	Username 	string `validate:"required,min=4,max=32,alphanum"`
	Password 	string `validate:"required,min=4,max=32,alphanum"`
	Name 		string `validate:"required,min=4,max=32,alpha"`
	Phone 		string `validate:"omitempty,max=13,numeric"`
}