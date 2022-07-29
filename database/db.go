package database

import (
	"log"

	"github.com/rafaelmgr12/Mark-URL/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnection() {
	connectionString := "host=database port=5432 user=root password=root dbname=root sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	DB.AutoMigrate(&models.URL{})

}
