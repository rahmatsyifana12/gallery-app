package models

import "time"

type User struct {
	ID 			uint64 `gorm:"primaryKey;autoIncrement"`
	Username 	string `gorm:"size:255;not null"`
	Password 	string `gorm:"size:255;not null"`
	Name 		string `gorm:"size:255;not null"`
	Phone 		string `gorm:"size:63"`
	CreatedAt 	time.Time `gorm:"not null"`
  	UpdatedAt 	time.Time `gorm:"not null"`
}