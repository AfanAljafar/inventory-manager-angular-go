package config

import (
	"backend-go/src/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DataBaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// ✅ Migrasi model User (tidak akan overwrite tabel existing)
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("❌ Failed to migrate user model: %v", err)
	}

	DB = db

	// ✅ Migrasi model Item (tidak akan overwrite tabel existing)
	err = db.AutoMigrate(&models.Item{})
	if err != nil {
		log.Fatalf("❌ Failed to migrate item model: %v", err)
	}

	fmt.Println("✅ Connected to database successfully!")
	return db
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database not initialized. Call DataBaseConnection() first.")
	}
	return DB
}
