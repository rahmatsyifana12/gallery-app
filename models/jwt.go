package models

import "github.com/golang-jwt/jwt/v4"

type JWTCustomClaims struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"size:255;not null"`
	jwt.RegisteredClaims
}
