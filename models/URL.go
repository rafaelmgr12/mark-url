package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	ID        string `json:"id" gorm:"primary_key"`
	URL       string `json:"url" gorm:"not null" validate:"required,url"`
	ShortURL  string `json:"short_url" validate:"required,url"`
	Hits      int    `json:"hits" gorm:"default:0"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func UrlValidated(url *URL) error {
	if err := validator.Validate(url); err != nil {
		return err
	}
	return nil
}
