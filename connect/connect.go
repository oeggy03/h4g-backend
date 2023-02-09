package connect

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/oeggy03/h4g-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// connecting to DB
var DB *gorm.DB

func Connect() {
	err := godotenv.Load()

	//if env file cannot be loaded
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Get DSN from .env file
	//from gorm documentation: connecting to MySQL
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = database
	//Once this is done, the table users will be created
	database.AutoMigrate(
		&models.User{},
	)
}
