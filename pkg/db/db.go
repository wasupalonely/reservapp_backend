package db

import (
	"fmt"
	"log"

	"github.com/wasupalonely/reservapp/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	var err error
	dsn := config.AppConfig.DatabaseUrl
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("Error opening database: %w", err)
	}

	log.Println("Database connection established!")
	return nil
}

// GetDB returns the database connection instance
func GetDB() *gorm.DB {
	return DB
}
