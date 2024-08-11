package main

import (
  "html/template"
  "log"
  "net/http"

  "github.com/your-project-name/config"
  "github.com/your-project-name/models"
)

func main() {
  // Database connection
  dsn := "user:password@tcp(your-db-host:3306)/your-db-name?charset=utf8mb4&parseTime=True&loc=Local"
  err := config.InitDB(dsn)
  if err != nil {
    log.Fatal("Failed to connect to database:", err)
  }

  // Migrate database schema (if needed)
  // ...

  // Load templates
  tmpl, err := template.ParseGlob("templates/*.html")
  if err != nil {
    log.Fatal("Error parsing templates:", err)
  }

  // Define routes
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    users := []models.User{}
    config.ConfigInstance.DB.Find(&users)

    err := tmpl.ExecuteTemplate(w, "index.html", struct{ Users []models.User }{Users: users})
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  })

  http.ListenAndServe(":8080", nil)
}
