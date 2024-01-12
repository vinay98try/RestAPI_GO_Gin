package config

import (
	"PracticeGDB/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// host=10.21.2.4
	dsn := "host=10.21.2.4 user=puser password=puser123 dbname=vidb port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print("error: - ", err.Error())
	}

	db.AutoMigrate(&models.User{})
	DB = db
}
