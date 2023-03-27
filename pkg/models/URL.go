package models

import (
	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	ID       string `gorm:"primary_key"`
	URL      string `json:"url" gorm:"not null" validate:"required,url"`
	ShortURL string `json:"short_url" validate:"required,url"`
	Public   bool   `json:"public" gorm:"default:false"`
	Hits     int    `json:"hits" gorm:"default:0"`
	UserID   string `json:"user_id" gorm:"default:null"`
}
