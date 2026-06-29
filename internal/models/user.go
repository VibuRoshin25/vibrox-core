package models

import "gorm.io/gorm"

// User is the user table model
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ID       int64  `json:"id" gorm:"primarykey"`
}
