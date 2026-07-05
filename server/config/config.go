package config

import (
	"log"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"StudentStats-backend-go/server/domain"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	log.Println("Database connection established")

	migrate()

	return DB
}

func migrate() {
	err := DB.AutoMigrate(
		&domain.Student{},
		// &domain.Subject{},
		// &domain.Grade{},
	)
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}
	log.Println("Database migration completed")
}
