# VykingProject

## Prerequisites 
### Docker Desktop
### MySQL Server
### Go v1.18 or higher

## Migrations done using go-migrate (https://github.com/golang-migrate/migrate) 
#### Created with command: `migrate create -ext sql -dir db/migrations -seq migration_name`

## MySQL server and database created using Docker Compose

## API routing done using mux (https://pkg.go.dev/github.com/gorilla/mux@v1.8.1)

## Swagger documentation done using Swag (https://github.com/swaggo/swag)

## Steps to start project
### 1. Start server and database with `docker-compose up -d`
### 2. Run the project using `go run .`
### 3. Verify migrations, database status, and server functionality
### 4. Open `http://localhost:8080/swagger/index.html#` to access Swagger UI and test API

## Migrations used in project
### 1. Create schemas for required tables
### 2. Create stored procedure for distributing prizes
### 3. Seed test data into the database
