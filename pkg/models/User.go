package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `gorm:"primary_key"`
	Username string `json:"username" gorm:"size:255;not null;unique"binding:"required" `
	Password string `json:"password" gorm:"size:255;not null;unique" binding:"required" sql:"password"`
	Email    string `json:"email" gorm:"size:255;not null;unique" binding:"required" `
}
