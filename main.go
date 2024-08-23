package main

import (
  "fmt"
  "os"
  "gorm.io/driver/mysql"
	"gorm.io/gorm"

  "github.com/joho/godotenv"

  "untitledgoproject/routes"
  "untitledgoproject/models"
)

var db *gorm.DB

func main() {
  err := godotenv.Load()
  if err != nil {
    panic(err)
  }

  // Get environment variables
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

  // Initialize database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})

  // Define routes & run
  routes.SetupRouter(db).Run(":8080")
}
