# Go Starter Kit

Go Starter Kit is a web application project built with Go. It leverages the following technologies:

- **GORM**: A popular ORM library for Go that provides a simple and intuitive way to interact with databases.
- **Gin**: A fast, lightweight, and flexible web framework for Go.

## Author

Derek Galbraith

## Installation

### 1. Install Go

```bash
sudo apt update
sudo apt install golang
which go
export PATH=$PATH:<path-to-your-go-bin>
source ~/.bashrc
go version
```

### 2. Install Required Packages

Open a terminal and run the following commands:

```bash
go get -u github.com/go-gorm/gorm
go get -u github.com/jinzhu/gorm/dialects/mysql
go get -u github.com/golang-migrate/migrate/v4/cmd/migrate
go get github.com/joho/godotenv
go get github.com/cosmtrek/air@latest
go get github.com/gin-gonic/gin
```

### 3. Clone the Repository

```bash
git clone https://github.com/derekg23/go-start-kit.git
cd go-start-kit
```

### 4. Install Dependencies

```bash
go mod tidy
```

### 5. Creating a MySQL Database

  Before running the application, you'll need to create a MySQL database.

#### 1. Install MySQL: If you don't already have MySQL installed, you can install it using the following commands:

  Linux (Ubuntu):

  ```bash
  sudo apt update
  sudo apt install mysql-server
  ```

  macOS (using Homebrew):

  ```bash
  brew install mysql
  ```

  Windows:

  Download and install MySQL from the [official website](https://www.postgresql.org/download/windows/).

#### 2. Create a Database:

  Start the MySQL service:
  
  Linux (Ubuntu):

  ```bash
  sudo service mysql start
  ```

  macOS:

  ```bash
  brew services start mysql
  ```

  Windows:

  Start the MySQL service from the Services management console or via MySQL Workbench.

  Switch to the MySQL interactive terminal (mysql):

  ```bash
  mysql -u root -p
  ```

  Create a new database:

  ```bash
  CREATE DATABASE <your-database-name>;
  ```

  (Optional) Create a new user and grant privileges:

  ```bash
  CREATE USER 'go_user'@'%' IDENTIFIED BY 'your_password';
  GRANT ALL PRIVILEGES ON <your-database-name>.* TO 'go_user'@'%';
  ```

### 5. Update environment file

  Copy .env.local to .env and update .env with your database information

  ```bash
  cp .env.local .env
  ```

### 6. Run Migrations

```bash
go run migrate.go
```

## Run Application

```bash
go run main.go
```


