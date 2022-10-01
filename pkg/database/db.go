package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rafaelmgr12/Mark-URL/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnection() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")

	connectionString := "host=" + DbHost + " user=" + DbUser + " password=" + DbPassword + " dbname=" + DbName + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	var models = []interface{}{&models.User{}, &models.URL{}}

	DB.AutoMigrate(models...)

}
