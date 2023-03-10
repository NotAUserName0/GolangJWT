package databases

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB) {
	
	dns := os.Getenv("USER") + ":" + os.Getenv("PASS") + "@tcp(127.0.0.1)/" + os.Getenv("DB") + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to BD")
	}

	log.Println("Database connected...")
	return DB
}