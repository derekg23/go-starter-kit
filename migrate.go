package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbURL := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Get the absolute path to the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory:", err)
	}

	// Construct the migration path using relative path
	migrationPath := filepath.Join(currentDir, "migrations")

	cmd := exec.Command("migrate", "-database", dbURL, "-path", migrationPath, "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error running migration: %v", err)
	}
}
