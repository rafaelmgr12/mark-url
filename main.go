package main

import (
	"log"

	"github.com/rafaelmgr12/Mark-URL/database"
	"github.com/rafaelmgr12/Mark-URL/routes"
)

func main() {
	log.Println("Starting application...")
	database.DatabaseConnection()
	log.Println("Database connection established")
	routes.HandleRequest()

}
