package config

import (
	"fmt"
	"os"

	"vibrox-core/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the gorm DB interface
var DB *gorm.DB

// Connect creates a DB connection
func Connect() {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		host, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err) //Panic is a built-in function that stops the ordinary flow of control and begins panicking.
	}
	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}

	DB = db
}
