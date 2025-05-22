package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"trendpulse-backend/internal/models"
)

var (
	db   *gorm.DB
	once sync.Once
)

// InitDB connects to the database and auto-migrates models
func InitDB() {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_USER", "postgres"),
			getEnv("DB_PASSWORD", "password"),
			getEnv("DB_NAME", "trendpulse"),
			getEnv("DB_PORT", "5432"),
		)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		// Auto-migrate the schema
		err = db.AutoMigrate(&models.User{}, &models.Article{})
		if err != nil {
			log.Fatalf("auto migration failed: %v", err)
		}

		log.Println("📦 Database connected and models migrated")
	})
}

// Get returns the GORM database instance
func Get() *gorm.DB {
	return db
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
