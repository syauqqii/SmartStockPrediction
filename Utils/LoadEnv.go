package Utils

import (
	"os"
	"fmt"
	"log"
	"github.com/joho/godotenv"
)

var (
	DB_CONN    string
	APP_HOST   string
	APP_PORT   string
	APP_CONF   string
	IS_DISPLAY string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal membaca file .env [LoadEnv]")
	}

	// Set Global Variable
	APP_HOST   = os.Getenv("APP_HOST")
	APP_PORT   = os.Getenv("APP_PORT")
	IS_DISPLAY = os.Getenv("IS_DISPLAY")

	// Lokal Variable
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	

	// Set Format (Database.go) -> root:@tcp(localhost:3306)/smartpredictstock
	DB_CONN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)

	// Set Format (Route.go) -> :8080
	APP_CONF = fmt.Sprintf(":%s", APP_PORT)
}