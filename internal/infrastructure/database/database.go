package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic("failed to load .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	if host == "" || port == "" || dbName == "" || user == "" || password == "" {
		panic("missing required environment variables")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable", host, user, password, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to PostgreSQL server!")
	}

	var count int
	db.Raw("SELECT COUNT(*) FROM pg_database WHERE datname = ?", dbName).Scan(&count)
	if count == 0 {
		if err := db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName)).Error; err != nil {
			panic(fmt.Sprintf("Failed to create database: %s", err))
		}
		log.Printf("Database %s created successfully", dbName)
	}

	db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database %s: %s", dbName, err))
	}

	return db
}
