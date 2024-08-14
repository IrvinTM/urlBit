package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB 

func init() {
	// Fetch environment variables 
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	if username == "" || password == "" || dbName == "" || dbHost == "" {
		log.Fatal("One or more environment variables (DB_USER, DB_PASS, DB_NAME, DB_HOST) are missing")
	}

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println("Connecting to the database...")

	// Retry
	var err error
	for i := 0; i < 5; i++ { 
		db, err = gorm.Open("postgres", dbUri)
		if err == nil {
			log.Println("Successfully connected to the database.")
			break
		}
		log.Printf("Failed to connect to database (attempt %d): %v", i+1, err)
		time.Sleep(5 * time.Second) 
	}

	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	db.Debug().AutoMigrate(&Account{}, &Url{}) // Database migration
}


func GetDB() *gorm.DB {
	return db
}
