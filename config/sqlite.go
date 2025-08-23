package config

import (
	"os"

	"github.com/amorimdotnetdevops-debug/go-user-control/schemas"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	// Initialize logger
	logger := GetLogger("sqlite")
	dbPath := "./db/access-control.db"

	// Check if the database file exists, if not create the directory
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Failed to create db directory: %v", err)
			return nil, err
		}
		// Create an empty database file
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Failed to create database file: %v", err)
			return nil, err
		}
		file.Close()
		logger.Info("Database file created.")
	}

	// Initialize SQLite connection
	logger.Info("Initializing SQLite database connection...")

	// Open SQLite database
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to connect to SQLite database: %v", err)
		return nil, err
	}

	// Auto-migrate the User schema
	err = db.AutoMigrate(&schemas.User{})
	if err != nil {
		logger.Errorf("SQLite autoMigration error: %v", err)
		return nil, err
	}

	logger.Info("Successfully connected to SQLite database.")
	return db, nil
}
