package models

type User struct {
	ID        string `json:"id" gorm:"primary_key" `
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	Email     string `json:"email" gorm:"unique"`
	TokenHash string `json:"token_hash"`
	CreatedAt string `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt string `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
