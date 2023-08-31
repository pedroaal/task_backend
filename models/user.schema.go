package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Mail      string
	Password  string
}

func (body *User) TableName() string {
	return "users"
}
