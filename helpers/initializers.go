package helpers

import (
	"GoJWT/databases"
	"GoJWT/models"
	"log"

	"github.com/joho/godotenv"
)

func LoadVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func MigrateDB(){
	db:= databases.Connection()
	db.AutoMigrate(&models.User{})
}