package main

import (
  "fmt"
  "html/template"
  "net/http"
  "os"
  "io/ioutil"
  "path/filepath"

  "github.com/joho/godotenv"
)

func router(url string) string {
  switch url {
    case "/":
      return "home.html"
    default:
      return "404.html"
  }
}

func handleInternalServerError(w http.ResponseWriter, err error) {
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func handleRoute(route string, tmpl *template.Template) {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    templateFilePath := filepath.Join("templates", "home.html")
    templateContent, err := ioutil.ReadFile(templateFilePath)
    handleInternalServerError(w, err)

    // Convert the file content to template.HTML
    data := struct {
      Template template.HTML
    } {
      Template: template.HTML(templateContent),
    }

    err = tmpl.ExecuteTemplate(w, "index.html", data)
    handleInternalServerError(w, err)
  })
}

func main() {
  err := godotenv.Load()
  if err != nil {
    panic(err)
  }

  appEnv := os.Getenv("APP_ENV")
  
  if appEnv == "local" {
    // Dislay starting message
    fmt.Println("\n\nYour website is now \x1b[1mrunning\x1b[0m")
    fmt.Printf("To view the website, visit \x1b[1mhttp://localhost:8080\x1b[0m\n\n\n")
  }

  // Load templates
  tmpl, err := template.ParseGlob("templates/*.html")
  if err != nil {
    panic(err)
  }

  // Define routes
  handleRoute("/", tmpl)

  // Serve static files from the "public" directory
  fs := http.FileServer(http.Dir("public"))
  http.Handle("/public/", http.StripPrefix("/public/", fs))

  http.ListenAndServe(":8080", nil)
}
