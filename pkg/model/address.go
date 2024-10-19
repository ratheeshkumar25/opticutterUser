package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model

	House  string `gorm:"not null"`
	Street string `gorm:"not null"`
	City   string `gorm:"not null"`
	ZIP    uint   `gorm:"not null"`
	State  string `gorm:"not null"`
	UserID uint   `gorm:"not null"`
}
