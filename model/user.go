package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"uniqueIndex,length:255"`
	FirstName string `gorm:"length:50,not null"`
	LastName  string `gorm:"length:100,not null"`
}
