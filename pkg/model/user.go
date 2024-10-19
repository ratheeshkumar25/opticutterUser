package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string  `gorm:"not null"`
	LastName  string  `gorm:"not null"`
	Phone     string  `gorm:"not null"`
	Email     string  `gorm:"not null; unique"`
	Password  string  `gorm:"not null"`
	Wallet    float64 `gorm:"default :0"`
	IsBlocked bool    `gorm:"default:false"`
}
