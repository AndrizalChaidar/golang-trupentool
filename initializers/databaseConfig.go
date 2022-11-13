package initializers

import (
	"golang-trupentool/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnect() {
	var err error
	dsn := os.Getenv("DB_CONFIG")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	DB.AutoMigrate(&models.Commander{}, &models.Troop{})
}
