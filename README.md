# Prerequisites

- MySQL (or your preferred database)

# Installation

### 1. Install Go:

```bash
sudo apt update
sudo apt install golang
which go
export PATH=$PATH:<path-to-your-go-bin>
source ~/.bashrc
go version
```

### 2. Install Required Packages:

Open a terminal and run the following commands:

```bash
go get -u github.com/go-gorm/gorm
go get -u github.com/jinzhu/gorm/dialects/mysql
go get -u github.com/golang-migrate/migrate/v4/cmd/migrate
go get github.com/joho/godotenv
go get github.com/cosmtrek/air@latest
```

### 3. Project Setup:

```bash
git clone https://github.com/your-username/your-project.git
cd your-project
```

### 4. Install Dependencies:

```bash
go mod tidy
```

#### 5. Create Environemnt File:

- Rename .env.local to .env
- Update .env file with db credentials

#### 6. Run Migrations:

```bash
go run migrate.go
```

# Run Application

#### 1. Run the Application:

```bash
go run main.go
```

# Run Migrations

#### 1. Run Migrations:

```bash
go run migrate.go
```


